package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) MoveDown(identifier uint) {
	var current seed.Seed
	r := s.mapper.Limit(1).Find(&current, identifier)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return
	}

	var below seed.Seed
	r = s.mapper.Where("position > ?", current.Position).
		Order("position ASC").Limit(1).Find(&below)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return
	}

	originalPosition := current.Position
	errors.PanicOnError(
		s.mapper.Model(&current).Update(
			"position",
			below.Position,
		).Error,
	)
	errors.PanicOnError(
		s.mapper.Model(&below).Update(
			"position",
			originalPosition,
		).Error,
	)
}
