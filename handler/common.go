package handler

import (
	"net/http"

	"michaelyusak/biaenergi-segment-generator.git/entity"

	"github.com/gin-gonic/gin"
)

type Common struct{}

func (h *Common) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, entity.Response{
		Code:    entity.CodeNotFound,
		Message: http.StatusText(http.StatusNotFound),
	})
}
