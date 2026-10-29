package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
	"time"
)

func (s *Service) TopAlerts(
	n int,
	start time.Time,
	end time.Time,
) ([]top.Record, error) {
	return s.store.Top(n, start, end)
}
