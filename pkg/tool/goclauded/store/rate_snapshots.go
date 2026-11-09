package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"
	"time"
)

func (s *Store) RateSnapshots(since time.Time) ([]rate_snapshot.Snapshot, error) {
	var result []rate_snapshot.Snapshot

	return result, s.database.Where("created_at > ?", since).Order(
		"created_at ASC",
	).Find(
		&result,
	).Error
}
