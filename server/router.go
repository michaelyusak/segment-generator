package server

import (
	"michaelyusak/biaenergi-segment-generator.git/handler"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	healthHandler *handler.Health
}

func createRouter(opt routerOpts) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
	)

	healthRouting(router, opt.healthHandler)

	return router
}

func healthRouting(r *gin.Engine, h *handler.Health) {
	r.GET("/health", h.Get)
}
