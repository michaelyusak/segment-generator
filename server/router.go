package server

import (
	"michaelyusak/biaenergi-segment-generator.git/handler"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	healthHandler *handler.Health
	canvasHandler *handler.Canvas
}

func createRouter(opt routerOpts) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
	)

	healthRouting(router, opt.healthHandler)
	canvasRouting(router, opt.canvasHandler)

	return router
}

func healthRouting(r *gin.Engine, h *handler.Health) {
	r.GET("/health", h.Get)
}

func canvasRouting(r *gin.Engine, h *handler.Canvas) {
	canvasGroup := r.Group("/v1/canvas")

	portGroup := canvasGroup.Group("/ports")
	portGroup.GET("", h.GetPorts)
	portGroup.GET("/:port_id", h.GetPort)
	portGroup.GET("/connections", h.GetConnections)

	nodeGroup := canvasGroup.Group("/nodes")
	nodeGroup.GET("", h.GetNodes)
	nodeGroup.GET("/:node_id", h.GetNode)
}
