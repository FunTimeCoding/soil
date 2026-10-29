package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"time"
)

func (s *Service) Snapshots(
	start time.Time,
	end time.Time,
	limit int,
) ([]snapshot.Snapshot, error) {
	return s.store.SnapshotsByTimeRange(start, end, limit)
}
