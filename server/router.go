package server

import (
	"michaelyusak/biaenergi-segment-generator.git/handler"
	"michaelyusak/biaenergi-segment-generator.git/middleware"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	healthHandler  *handler.Health
	canvasHandler  *handler.Canvas
	segmentHandler *handler.Segment
	commonHandler  handler.Common
}

func createRouter(opt routerOpts) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

	router.NoRoute(opt.commonHandler.NotFound)

	healthRouting(router, opt.healthHandler)
	canvasRouting(router, opt.canvasHandler)
	segmentRouting(router, opt.segmentHandler)

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

func segmentRouting(r *gin.Engine, h *handler.Segment) {
	segmentGroup := r.Group("/v1/segments")
	segmentGroup.GET("", h.GetSegments)
}
