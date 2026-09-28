package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"time"
)

func (s *Store) SweepPlacements(
	source string,
	before time.Time,
) (int64, error) {
	result := s.mapper.Where(
		constant.SweepCondition,
		source,
		before,
	).Delete(placement.Stub())

	return result.RowsAffected, result.Error
}
