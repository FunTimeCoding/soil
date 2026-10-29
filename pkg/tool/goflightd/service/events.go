package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"time"
)

func (s *Service) Events(
	start time.Time,
	end time.Time,
	limit int,
) ([]event.Event, error) {
	return s.store.EventsByTimeRange(start, end, limit)
}
