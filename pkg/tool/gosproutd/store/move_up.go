package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) MoveUp(identifier uint) {
	var current seed.Seed
	r := s.mapper.Limit(1).Find(&current, identifier)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return
	}

	var above seed.Seed
	r = s.mapper.Where("position < ?", current.Position).
		Order("position DESC").Limit(1).Find(&above)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return
	}

	originalPosition := current.Position
	errors.PanicOnError(
		s.mapper.Model(&current).Update(
			"position",
			above.Position,
		).Error,
	)
	errors.PanicOnError(
		s.mapper.Model(&above).Update(
			"position",
			originalPosition,
		).Error,
	)
}
