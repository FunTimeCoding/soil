package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
	"time"
)

func (s *Store) FableSnapshotsSince(since time.Time) ([]fable_snapshot.Snapshot, error) {
	var result []fable_snapshot.Snapshot

	return result, s.database.Where("created_at > ?", since).Order(
		"created_at ASC",
	).Find(
		&result,
	).Error
}
