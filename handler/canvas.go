package handler

import (
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"michaelyusak/biaenergi-segment-generator.git/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Canvas struct {
	canvasService service.Canvas
}

func NewCanvas(canvasService service.Canvas) *Canvas {
	return &Canvas{
		canvasService: canvasService,
	}
}

func (h *Canvas) GetPorts(ctx *gin.Context) {
	ports, err := h.canvasService.GetPorts(ctx.Request.Context())
	if err != nil {
		logrus.WithError(err).Error("[handler][Canvas][GetPorts] failed to get ports")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    ports,
	})
}

func (h *Canvas) GetPort(ctx *gin.Context) {
	portID := ctx.Param("port_id")
	if portID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.Response{
			Code:    entity.CodeBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	port, err := h.canvasService.GetPort(ctx.Request.Context(), portID)
	if err != nil {
		logrus.WithError(err).WithField("port_id", portID).
			Error("[handler][Canvas][GetPort] failed to get port")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if port == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, entity.Response{
			Code:    entity.CodeNotFound,
			Message: http.StatusText(http.StatusNotFound),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    port,
	})
}

func (h *Canvas) GetNodes(ctx *gin.Context) {
	nodes, err := h.canvasService.GetNodes(ctx.Request.Context())
	if err != nil {
		logrus.WithError(err).Error("[handler][Canvas][GetNodes] failed to get nodes")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    nodes,
	})
}

func (h *Canvas) GetNode(ctx *gin.Context) {
	nodeIDStr := ctx.Param("node_id")
	if nodeIDStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.Response{
			Code:    entity.CodeBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	nodeID, err := strconv.ParseInt(nodeIDStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.Response{
			Code:    entity.CodeBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	node, err := h.canvasService.GetNode(ctx.Request.Context(), nodeID)
	if err != nil {
		logrus.WithError(err).WithField("node_id", nodeID).
			Error("[handler][Canvas][GetNode] failed to get node")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if node == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, entity.Response{
			Code:    entity.CodeNotFound,
			Message: http.StatusText(http.StatusNotFound),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    node,
	})
}

func (h *Canvas) GetConnections(ctx *gin.Context) {
	connections, err := h.canvasService.GetConnections(ctx.Request.Context())
	if err != nil {
		logrus.WithError(err).Error("[handler][Canvas][GetConnections] failed to get connections")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    connections,
	})
}
