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
	browser.OpenWeek()
	time.Sleep(time.Second)
	s.RecordUsagePage(browser.ReadUsage())
}
