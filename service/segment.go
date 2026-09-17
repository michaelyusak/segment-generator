package service

import (
	"context"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"michaelyusak/biaenergi-segment-generator.git/repository"
)

type segmentService struct {
	portRepository repository.Port
}

func NewSegmentService(portRepository repository.Port) *segmentService {
	return &segmentService{
		portRepository: portRepository,
	}
}

func (s *segmentService) GetSegments(ctx context.Context) ([]entity.Segment, error) {
	connections, err := s.portRepository.GetAllConnections(ctx)
	if err != nil {
		return nil, fmt.Errorf("[service][segmentService][GetSegmentCounts] failed to get all connections: %w", err)
	}

	heads, err := s.portRepository.GetSegmentHeads(ctx)
	if err != nil {
		return nil, fmt.Errorf("[service][segmentService][GetSegmentCounts] failed to get segment heads: %w", err)
	}

	sourcesByPortID := map[string][]entity.Port{}

	for _, conn := range connections {
		sourcesByPortID[conn.Target.ID] = append(sourcesByPortID[conn.Target.ID], conn.Source)
	}

	segments := []entity.Segment{}

	for _, head := range heads {
		segments = append(segments, s.buildSegments(head, sourcesByPortID)...)
	}

	return segments, nil

}

func (s *segmentService) buildSegments(head entity.Port, sourcesByPortID map[string][]entity.Port) []entity.Segment {
	segments := []entity.Segment{}
	queue := []entity.Port{}
	queue = append(queue, head)

	var dfs func(segment *entity.Segment, target entity.Port)
	dfs = func(segment *entity.Segment, target entity.Port) {
		first := segment == nil
		if first {
			segment = &entity.Segment{
				Target:  &target,
				Sources: []*entity.Port{},
			}
		}

		sources, ok := sourcesByPortID[target.ID]
		if !ok || len(sources) == 0 {
			return
		}

		for _, source := range sources {
			if source.Value == nil {
				dfs(segment, source)

				continue
			}

			segment.Sources = append(segment.Sources, &source)
			queue = append(queue, sourcesByPortID[source.ID]...)
		}

		if first && len(segment.Sources) > 0 {
			segment.Finalise()
			segments = append(segments, *segment)
		}
	}

	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]

		dfs(nil, h)
	}

	return segments
}
