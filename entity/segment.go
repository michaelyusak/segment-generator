package entity

import (
	"fmt"
	"strings"
)

type Segment struct {
	Name    string  `json:"name"`
	Value   int64   `json:"value"`
	Target  *Port   `json:"target"`
	Sources []*Port `json:"sources"`
}

func (s *Segment) Finalise() {
	sourceIDs := make([]string, 0, len(s.Sources))

	var sumSourceValues int64

	for _, source := range s.Sources {
		sourceIDs = append(sourceIDs, source.ID)
		sumSourceValues += *source.Value
	}

	s.Value = *s.Target.Value - sumSourceValues

	if len(s.Sources) == 1 {
		s.Name = fmt.Sprintf("%s - %s", s.Target.ID, sourceIDs[0])
	} else {
		s.Name = fmt.Sprintf("%s - (%s)", s.Target.ID, strings.Join(sourceIDs, " + "))
	}
}
