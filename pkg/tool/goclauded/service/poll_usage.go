package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site"
	"time"
)

func (s *Service) PollUsage() {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Structured(
				"usage poll failed",
				"error",
				r,
			)
		}
	}()
	browser := site.New()
	browser.ClickRefresh()
	time.Sleep(2 * time.Second)
	result := browser.ReadUsage()

	if result == nil {
		return
	}

	s.usageMutex.Lock()
	s.usage = result
	s.usageMutex.Unlock()
	s.store.SaveUsageSnapshot(
		result.SessionPercent,
		result.WeeklyAllPercent,
		result.WeeklySonnetPercent,
		result.CreditPercent,
	)
	s.store.TrimUsageSnapshots()
	s.notifier.Notify()
}
