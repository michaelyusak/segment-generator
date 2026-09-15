package server

import (
	"context"
	"michaelyusak/biaenergi-segment-generator.git/config"
	"michaelyusak/biaenergi-segment-generator.git/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func Init() {
	config, err := config.Init()
	if err != nil {
		logrus.WithError(err).Fatal("config init failed")
	}

	healthHandler := handler.NewHealth()

	router := createRouter(routerOpts{
		healthHandler: healthHandler,
	})

	srv := http.Server{
		Handler: router,
		Addr:    config.Port,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("listen")
		}
	}()

	quit := make(chan os.Signal, 10)

	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.WithField("graceful_period", time.Duration(config.GracefulPeriod).String()).Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GracefulPeriod))
	defer cancel()

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Server shutdown")
	}

	logrus.Info("Server exiting")
}
