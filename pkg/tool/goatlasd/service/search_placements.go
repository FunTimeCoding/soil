package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

func (s *Service) SearchPlacements(name string) ([]*placement.Placement, error) {
	return s.store.SearchPlacements(name)
}
