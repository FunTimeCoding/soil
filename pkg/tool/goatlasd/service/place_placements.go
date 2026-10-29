package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

func (s *Service) PlacePlacements(
	kind string,
	name string,
) ([]*placement.Placement, error) {
	return s.store.PlacePlacements(kind, name)
}
