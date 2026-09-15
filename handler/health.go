package handler

import (
	"net/http"

	"michaelyusak/biaenergi-segment-generator.git/entity"

	"github.com/gin-gonic/gin"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Get(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		entity.Response{
			Code:    "SUCCESS",
			Message: "ok",
		})
}
