package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store/result"
)

func (s *Store) Places() ([]*result.Place, error) {
	var v []*result.Place

	return v, s.mapper.Model(placement.Stub()).Select(
		constant.PlaceSelect,
	).Group(constant.PlaceGroup).Order(constant.PlaceOrder).Scan(&v).Error
}
