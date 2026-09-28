package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
)

func (s *Store) PlaceSightings(
	kind string,
	name string,
) ([]*sighting.Sighting, error) {
	var result []*sighting.Sighting

	return result, s.mapper.Order(constant.SightingOrder).Where(
		constant.PlaceKindCondition,
		kind,
	).Where(constant.PlaceCondition, name).Find(&result).Error
}
