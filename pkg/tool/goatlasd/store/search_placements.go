package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
)

func (s *Store) SearchPlacements(name string) ([]*placement.Placement, error) {
	var result []*placement.Placement
	m := s.mapper.Order(constant.PlacementOrder)

	if name != "" {
		m = m.Where(
			constant.NameCondition,
			fmt.Sprintf(constant.NamePattern, name),
		)
	}

	return result, m.Find(&result).Error
}
