package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"
)

func (s *Store) Frames(session string) []*frame.Frame {
	var result []*frame.Frame
	errors.PanicOnError(
		s.mapper.Where(
			"session = ?",
			session,
		).Order(constant.IdentifierColumn).Find(&result).Error,
	)

	return result
}
