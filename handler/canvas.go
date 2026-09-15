package handler

import (
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"michaelyusak/biaenergi-segment-generator.git/service"
	"net/http"

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
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    ports,
	})
}
