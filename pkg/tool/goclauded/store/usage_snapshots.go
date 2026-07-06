package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/usage_snapshot"
	"time"
)

func (s *Store) UsageSnapshots(since time.Duration) ([]usage_snapshot.Snapshot, error) {
	var result []usage_snapshot.Snapshot

	return result, s.database.Where(
		"created_at > ?",
		s.clock().Add(-since),
	).Order("created_at ASC").Find(&result).Error
}
