package usage

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
	"time"
)

func (s *Service) PollUsage() []error {
	if !chromium.ListeningEnvironment() {
		s.logger.Structured("usage poll skipped", "reason", "browser")

		return nil
	}

	browser := site.Attach(s.browser)
	browser.ClickRefresh()
	time.Sleep(2 * time.Second)
	browser.OpenWeek()
	time.Sleep(time.Second)

	if e := s.RecordUsagePage(browser.ReadUsage()); e != nil {
		return []error{e}
	}

	return nil
}
