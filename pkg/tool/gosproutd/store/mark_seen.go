package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) MarkSeen(session string) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"session = ? AND answered_at IS NOT NULL AND seen_at IS NULL",
			session,
		).Update("seen_at", s.clock()).Error,
	)
}
