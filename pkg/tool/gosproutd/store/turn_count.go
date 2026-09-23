package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
)

func (s *Store) TurnCount(
	identifier uint,
	author constant.Author,
) int64 {
	var result int64
	errors.PanicOnError(
		s.mapper.Model(turn.Stub()).Where(
			"decision_identifier = ? AND author = ?",
			identifier,
			author,
		).Count(&result).Error,
	)

	return result
}
