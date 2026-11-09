package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
)

func (s *Service) topWindow(w *window.Window) *window.Window {
	return w.Bound(s.clock(), constant.TopSpan)
}
