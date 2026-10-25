package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"time"
)

func (s *Store) EventsByTimeRange(
	start time.Time,
	end time.Time,
	limit int,
) ([]event.Event, error) {
	var result []event.Event
	e := s.database.
		Where("time >= ? AND time <= ?", start, end).
		Order("time").
		Limit(limit).
		Find(&result).Error

	return result, e
}
