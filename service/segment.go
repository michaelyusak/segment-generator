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
	pathMap, err := s.portRepository.GetPaths(ctx)
	if err != nil {
		return nil, fmt.Errorf("[service][segmentService][GetSegments] failed to get paths: %w", err)
	}

	segments := []entity.Segment{}

	pending := pathMap

	for len(pending) > 0 {
		next := make(map[string][][]entity.Port)

		for _, paths := range pending {
			segment := entity.Segment{
				Target: &paths[0][0],
			}

			visited := map[string]bool{}

			for _, path := range paths {
				for i := 1; i < len(path); i++ {
					port := path[i]

					if port.Value == nil {
						continue
					}

					if !visited[port.ID] {
						visited[port.ID] = true
						segment.Sources = append(segment.Sources, &port)
					}

					next[port.ID] = append(next[port.ID], path[i:])
					break
				}
			}

			if len(segment.Sources) < 1 {
				continue
			}

			segment.Finalise()
			segments = append(segments, segment)
		}

		pending = next
	}

	return segments, nil

}
