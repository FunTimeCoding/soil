package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) UnpulsedAnswers(session string) []*decision.Decision {
	var result []*decision.Decision
	errors.PanicOnError(
		s.mapper.Where(
			"session = ? AND answered_at IS NOT NULL AND seen_at IS NULL AND pulsed_at IS NULL",
			session,
		).Order(constant.IdentifierColumn).Find(&result).Error,
	)

	return result
}
