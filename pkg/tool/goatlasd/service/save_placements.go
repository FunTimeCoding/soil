package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (s *Service) savePlacements(
	source string,
	v []*placement.Placement,
) (int64, error) {
	for _, p := range v {
		if e := s.store.SavePlacement(p); e != nil {
			return 0, e
		}
	}

	return s.store.SweepPlacements(source, time.Now().Add(-s.retention))
}
