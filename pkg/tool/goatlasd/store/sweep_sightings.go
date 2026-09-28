package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"time"
)

func (s *Store) SweepSightings(
	source string,
	before time.Time,
) (int64, error) {
	result := s.mapper.Where(
		constant.SweepCondition,
		source,
		before,
	).Delete(sighting.Stub())

	return result.RowsAffected, result.Error
}
