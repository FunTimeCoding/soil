package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) UnseenAnswers(session string) []*decision.Decision {
	var result []*decision.Decision
	errors.PanicOnError(
		s.mapper.Preload("Frames").Where(
			"session = ? AND answered_at IS NOT NULL AND seen_at IS NULL",
			session,
		).Order(constant.IdentifierColumn).Find(&result).Error,
	)

	return result
}
