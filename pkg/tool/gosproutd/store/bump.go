package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"gorm.io/gorm"
)

func (s *Store) Bump(identifier uint) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"identifier = ?",
			identifier,
		).UpdateColumn("bumped", gorm.Expr("bumped + 1")).Error,
	)
}
