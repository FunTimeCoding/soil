package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Store) CleanupQueue(cutoff time.Time) int64 {
	result := s.database.Where(
		"consumed = ? AND created_at < ?",
		true,
		cutoff,
	).Delete(queue.Stub())

	return result.RowsAffected
}
