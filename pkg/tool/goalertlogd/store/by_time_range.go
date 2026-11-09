package store

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func (s *Store) ByTimeRange(w *window.Window) ([]record.Record, error) {
	var result []record.Record
	q := s.database

	if start, started := w.Start(); started {
		q = q.Where("started_at >= ?", start)
	}

	if end, ended := w.End(); ended {
		q = q.Where("started_at < ?", end)
	}

	return result, q.Order("started_at").Find(&result).Error
}
