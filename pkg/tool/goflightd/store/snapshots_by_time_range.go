package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (s *Store) SnapshotsByTimeRange(
	start time.Time,
	end time.Time,
	limit int,
) ([]snapshot.Snapshot, error) {
	var result []snapshot.Snapshot
	e := s.database.
		Where("time >= ? AND time <= ?", start, end).
		Order("time").
		Limit(limit).
		Find(&result).Error

	return result, e
}
