package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (s *Store) UpdateCompletionSummary(
	sessionIdentifier string,
	topic string,
	message string,
) {
	errors.PanicOnError(
		s.database.Model(completion.Stub()).
			Where(
				"session_identifier = ? AND topic = ?",
				sessionIdentifier,
				topic,
			).
			Update(constant.SummaryColumn, message).Error,
	)
}
