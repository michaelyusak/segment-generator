package service

import (
	"context"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"michaelyusak/biaenergi-segment-generator.git/repository"
)

type canvasService struct {
	portRepository repository.Port
}

func NewCanvasService(portRepository repository.Port) *canvasService {
	return &canvasService{
		portRepository: portRepository,
	}
}

func (s *canvasService) GetPorts(ctx context.Context) ([]entity.Port, error) {
	ports, err := s.portRepository.GetPorts(ctx)
	if err != nil {
		return nil, fmt.Errorf("[service][canvasService] failed to get ports: %w", err)
	}

	return ports, nil
}
