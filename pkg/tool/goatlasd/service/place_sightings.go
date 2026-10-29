package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"

func (s *Service) PlaceSightings(
	kind string,
	name string,
) ([]*sighting.Sighting, error) {
	return s.store.PlaceSightings(kind, name)
}
