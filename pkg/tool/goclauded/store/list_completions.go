package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (s *Store) ListCompletions() []completion.Completion {
	var result []completion.Completion
	errors.PanicOnError(
		s.database.Order("created_at ASC").Find(&result).Error,
	)

	return result
}
