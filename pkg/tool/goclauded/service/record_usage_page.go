package service

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) RecordUsagePage(result *page.Usage) {
	if result == nil {
		if !s.usageUnreadable {
			s.usageUnreadable = true
			s.reporter.CaptureException(
				errors.New(constant.UsagePageUnreadable),
			)
		}

		return
	}

	if s.usageUnreadable {
		s.usageUnreadable = false
		s.logger.Structured("usage page readable again")
	}

	if !result.FableSeen {
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
