package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"

func (s *Service) MatchingSightings(
	place string,
	unclaimed bool,
) ([]*sighting.Sighting, error) {
	return s.store.MatchingSightings(place, unclaimed)
}
