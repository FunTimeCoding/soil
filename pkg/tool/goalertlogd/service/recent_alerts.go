package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"time"
)

func (s *Service) RecentAlerts(
	start time.Time,
	end time.Time,
) ([]record.Record, error) {
	return s.store.ByTimeRange(start, end)
}
