package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
	"time"
)

func (s *Service) PollUsage() {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Structured("usage poll failed", "error", r)
		}
	}()
	browser := site.New()
	browser.ClickRefresh()
	time.Sleep(2 * time.Second)
	result := browser.ReadUsage()

	if result == nil {
		return
	}

	if e := s.recordFable(result.FablePercent, result.FableReset); e != nil {
		s.logger.Structured("fable snapshot failed", "error", e)
	}

	s.notifier.Notify()
}
