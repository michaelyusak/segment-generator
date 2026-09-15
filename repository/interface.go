package repository

import (
	"context"
	"michaelyusak/biaenergi-segment-generator.git/entity"
)

type Node interface {
	GetNodes(ctx context.Context) ([]entity.Node, error)
	GetNode(ctx context.Context, nodeID int64) (*entity.Node, error)
}

type Port interface {
	GetPorts(ctx context.Context) ([]entity.Port, error)
	GetPort(ctx context.Context, portID string) (*entity.Port, error)
}
