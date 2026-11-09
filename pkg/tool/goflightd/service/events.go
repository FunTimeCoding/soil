package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
)

func (s *Service) Events(
	w *window.Window,
	limit int,
) ([]event.Event, error) {
	return s.store.Events(w.Bound(s.clock(), constant.QuerySpan), limit)
}
