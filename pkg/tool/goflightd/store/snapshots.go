package store

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
)

func (s *Store) Snapshots(
	w *window.Window,
	limit int,
) ([]snapshot.Snapshot, error) {
	var result []snapshot.Snapshot
	q := s.database

	if start, started := w.Start(); started {
		q = q.Where("time >= ?", start)
	}

	if end, ended := w.End(); ended {
		q = q.Where("time <= ?", end)
	}

	e := q.Order("time").Limit(limit).Find(&result).Error

	return result, e
}
