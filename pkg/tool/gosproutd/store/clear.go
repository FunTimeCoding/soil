package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) Clear(
	identifier uint,
	line string,
	resolution constant.Resolution,
) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"identifier = ?",
			identifier,
		).Updates(
			map[string]any{"clear_line": line, "resolution": resolution},
		).Error,
	)
}
