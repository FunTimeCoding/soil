package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/cruise"
)

func (s *Store) MarkPromoted(session string) {
	errors.PanicOnError(
		s.mapper.Model(cruise.Stub()).Where(
			"session = ?",
			session,
		).Update("promoted_at", s.clock()).Error,
	)
}
