package handler

import (
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"michaelyusak/biaenergi-segment-generator.git/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Segment struct {
	segmentService service.Segment
}

func NewSegment(segmentService service.Segment) *Segment {
	return &Segment{
		segmentService: segmentService,
	}
}

func (h *Segment) GetSegments(ctx *gin.Context) {
	segments, err := h.segmentService.GetSegments(ctx.Request.Context())
	if err != nil {
		logrus.WithError(err).Error("[handler][Segment][GetSegments] failed to get segments")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.Response{
			Code:    entity.CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.Response{
		Code:    entity.CodeSuccess,
		Message: http.StatusText(http.StatusOK),
		Data: gin.H{
			"count":    len(segments),
			"segments": segments,
		},
	})
}
