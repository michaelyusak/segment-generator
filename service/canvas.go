package service

import (
	"context"
	"errors"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/apperror"
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
		return nil, fmt.Errorf("[service][canvasService][GetPorts] failed to get ports: %w", err)
	}

	return ports, nil
}


func (s *canvasService) GetPort(ctx context.Context, portID string) (*entity.Port, error) {
	port, err := s.portRepository.GetPort(ctx, portID)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("[service][canvasService][GetPort] failed to get port: %w [port_id: %s]", err, portID)
	}

	return port, nil
}