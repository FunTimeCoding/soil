package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (s *Store) MatchingPlacements(
	place string,
	source string,
	packageName string,
) ([]*placement.Placement, error) {
	var result []*placement.Placement
	m := s.mapper.Order(constant.PlacementOrder)

	if place != "" {
		m = m.Where(constant.PlaceCondition, place)
	}

	if source != "" {
		m = m.Where(constant.SourceCondition, source)
	}

	if packageName != "" {
		m = m.Where(constant.PackageCondition, packageName)
	}

	return result, m.Find(&result).Error
}
