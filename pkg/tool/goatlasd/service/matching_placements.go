package service

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

func (s *Service) MatchingPlacements(
	place string,
	source string,
	packageName string,
) ([]*placement.Placement, error) {
	return s.store.MatchingPlacements(place, source, packageName)
}
