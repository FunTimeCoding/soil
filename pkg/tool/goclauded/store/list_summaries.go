package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"
)

func (s *Store) ListSummaries() []summary.Summary {
	var result []summary.Summary
	errors.PanicOnError(s.database.Find(&result).Error)

	return result
}
