package service

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
	"time"
)

func (s *Service) PollUsage() {
	if !chromium.ListeningEnvironment() {
		s.logger.Structured("usage poll skipped", "reason", "browser")

		return
	}

	browser := site.Attach(s.usageBrowser())
	browser.ClickRefresh()
	time.Sleep(2 * time.Second)
	result := browser.ReadUsage()

	if result == nil {
		return
	}

	if e := s.RecordFable(
		result.FablePercent,
		result.FableReset,
		nil,
	); e != nil {
		s.logger.Structured("fable snapshot failed", "error", e)
	}

	s.notifier.Notify()
}
