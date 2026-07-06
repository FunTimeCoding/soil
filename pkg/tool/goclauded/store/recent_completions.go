package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
	"time"
)

func (s *Store) RecentCompletions() ([]completion.Completion, error) {
	var result []completion.Completion

	return result, s.database.Where(
		"created_at > ?",
		s.clock().Add(-24*time.Hour),
	).Order("created_at DESC").Find(&result).Error
}
