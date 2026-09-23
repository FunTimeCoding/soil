package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) Dismiss(identifier uint, state constant.State) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"identifier = ?",
			identifier,
		).Updates(
			map[string]any{"state": state, "answered_at": s.clock()},
		).Error,
	)
}
