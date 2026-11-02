package usage

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) RecordUsagePage(result *page.Usage) error {
	if result == nil {
		if !s.usageUnreadable {
			s.usageUnreadable = true
			s.reporter.CaptureException(
				errors.New(constant.UsagePageUnreadable),
			)
		}

		return nil
	}

	if s.usageUnreadable {
		s.usageUnreadable = false
		s.logger.Structured("usage page readable again")
	}

	if !result.FableSeen {
		return nil
	}

	e := s.RecordFable(result.FablePercent, result.FableReset, nil)
	s.notifier.Notify()

	return e
}
