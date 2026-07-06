package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"gorm.io/gorm"
)

func (s *Store) SetPosition(
	identifier uint,
	target int,
) {
	var current seed.Seed
	r := s.mapper.Limit(1).Find(&current, identifier)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return
	}

	origin := current.Position

	if target == origin {
		return
	}

	if target < origin {
		errors.PanicOnError(
			s.mapper.Model(seed.Stub()).
				Where("position >= ? AND position < ?", target, origin).
				Update("position", gorm.Expr("position + 1")).Error,
		)
	} else {
		errors.PanicOnError(
			s.mapper.Model(seed.Stub()).
				Where("position > ? AND position <= ?", origin, target).
				Update("position", gorm.Expr("position - 1")).Error,
		)
	}

	errors.PanicOnError(
		s.mapper.Model(&current).Update("position", target).Error,
	)
}
