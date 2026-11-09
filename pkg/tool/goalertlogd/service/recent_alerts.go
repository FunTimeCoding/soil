package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func (s *Service) RecentAlerts(w *window.Window) ([]record.Record, error) {
	return s.store.ByTimeRange(s.RecentWindow(w))
}
