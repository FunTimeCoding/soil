package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"time"
)

func (s *Service) spanStart(span time.Duration) time.Time {
	result, _ := window.Whole().Bound(s.clock(), span).Start()

	return result
}
