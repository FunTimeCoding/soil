package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) MarkPulsed(session string) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"session = ? AND answered_at IS NOT NULL AND seen_at IS NULL AND pulsed_at IS NULL",
			session,
		).Update("pulsed_at", s.clock()).Error,
	)
}
