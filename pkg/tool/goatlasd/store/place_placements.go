package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (s *Store) PlacePlacements(
	kind string,
	name string,
) ([]*placement.Placement, error) {
	var result []*placement.Placement

	return result, s.mapper.Order(constant.PlacementOrder).Where(
		constant.PlaceKindCondition,
		kind,
	).Where(constant.PlaceCondition, name).Find(&result).Error
}
