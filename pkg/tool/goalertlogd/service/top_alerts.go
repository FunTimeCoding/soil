package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
)

func (s *Service) TopAlerts(
	n int,
	w *window.Window,
) ([]top.Record, error) {
	return s.store.Top(n, s.topWindow(w))
}
