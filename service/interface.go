package service

import (
	"context"
	"michaelyusak/biaenergi-segment-generator.git/entity"
)

type Canvas interface {
	GetPorts(ctx context.Context) ([]entity.Port, error)
}
