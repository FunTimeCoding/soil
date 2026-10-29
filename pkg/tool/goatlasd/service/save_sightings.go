package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"time"
)

func (s *Service) saveSightings(
	source string,
	v []*sighting.Sighting,
) (int64, error) {
	for _, p := range v {
		if e := s.store.SaveSighting(p); e != nil {
			return 0, e
		}
	}

	return s.store.SweepSightings(source, time.Now().Add(-s.retention))
}
