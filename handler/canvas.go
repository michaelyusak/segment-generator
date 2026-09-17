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

// Get canvas ports
// @Summary Get canvas ports
// @Description Returns all canvas ports
// @Tags canvas
// @Produce json
// @Success 200 {object} entity.Response{data=[]entity.Port}
// @Failure 500 {object} entity.Response
// @Router /v1/canvas/ports [get]
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

// Get canvas port detail
// @Summary Get canvas port detail
// @Description Returns canvas port detail
// @Tags canvas
// @Produce json
// @Param port_id path string true "Port ID"
// @Success 200 {object} entity.Response{data=entity.Port}
// @Failure 400 {object} entity.Response
// @Failure 404 {object} entity.Response
// @Failure 500 {object} entity.Response
// @Router /v1/canvas/ports/{port_id} [get]
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

// Get canvas nodes
// @Summary Get canvas nodes
// @Description Returns all canvas nodes
// @Tags canvas
// @Produce json
// @Success 200 {object} entity.Response{data=[]entity.Node}
// @Failure 500 {object} entity.Response
// @Router /v1/canvas/nodes [get]
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

// Get canvas node detail
// @Summary Get canvas node detail
// @Description Returns canvas node detail
// @Tags canvas
// @Produce json
// @Param node_id path int64 true "Node ID"
// @Success 200 {object} entity.Response{data=entity.Node}
// @Failure 400 {object} entity.Response
// @Failure 404 {object} entity.Response
// @Failure 500 {object} entity.Response
// @Router /v1/canvas/nodes/{node_id} [get]
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

// Get canvas port connections
// @Summary Get canvas port connections
// @Description Returns all canvas port connections
// @Tags canvas
// @Produce json
// @Success 200 {object} entity.Response{data=[]entity.PortConnection}
// @Failure 500 {object} entity.Response
// @Router /v1/canvas/ports/connections [get]
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
