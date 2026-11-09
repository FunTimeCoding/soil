package store

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
)

func (s *Store) Events(
	w *window.Window,
	limit int,
) ([]event.Event, error) {
	var result []event.Event
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
