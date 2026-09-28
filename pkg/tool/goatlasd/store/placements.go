package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (s *Store) Placements() ([]*placement.Placement, error) {
	var result []*placement.Placement

	return result, s.mapper.Order(constant.PlacementOrder).Find(&result).Error
}
