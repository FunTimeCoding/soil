package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
)

func (s *Store) MatchingSightings(
	place string,
	unclaimed bool,
) ([]*sighting.Sighting, error) {
	var result []*sighting.Sighting
	m := s.mapper.Order(constant.SightingOrder)

	if place != "" {
		m = m.Where(constant.PlaceCondition, place)
	}

	if unclaimed {
		m = m.Where(constant.UnclaimedCondition)
	}

	return result, m.Find(&result).Error
}
