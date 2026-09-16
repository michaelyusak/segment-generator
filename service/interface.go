package service

import (
	"context"
	"michaelyusak/biaenergi-segment-generator.git/entity"
)

type Canvas interface {
	// ports
	GetPorts(ctx context.Context) ([]entity.Port, error)
	GetPort(ctx context.Context, portID string) (*entity.Port, error)

	// nodes
	GetNodes(ctx context.Context) ([]entity.Node, error)
	GetNode(ctx context.Context, nodeID int64) (*entity.Node, error)

	// connections
	GetConnections(ctx context.Context) ([]entity.PortConnection, error)
}

type Segment interface {
	GetSegments(ctx context.Context) ([]entity.Segment, error)
}
