package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Store) CleanupQueue(cutoff time.Time) int64 {
	result := s.database.Where(
		"consumed_at IS NOT NULL AND created_at < ?",
		cutoff,
	).Delete(
		queue.Stub(),
	)

	return result.RowsAffected
}
