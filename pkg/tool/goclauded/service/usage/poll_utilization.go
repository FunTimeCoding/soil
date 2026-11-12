package usage

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/usage_result"
)

func (s *Service) PollUtilization(x context.Context) []error {
	c := s.credential()

	if c == nil {
		s.logger.Structured("utilization poll skipped", "reason", "credential")

		return nil
	}

	if c.Expired(s.clock()) {
		s.logger.Structured(
			"utilization poll skipped",
			"reason",
			"expired",
			"expiredAt",
			c.ExpiresAt,
		)

		return nil
	}

	var result *usage_result.Result

	if e := s.utilizationGuard.Attempt(
		x,
		func(y context.Context) error {
			var f error
			result, f = s.utilization.Read(y, c.AccessToken)

			return f
		},
	); e != nil {
		return []error{e}
	}
	var failures []error

	if f := s.RecordRateLimits(
		result.SessionPercent,
		result.WeeklyPercent,
		result.SessionReset,
		result.WeeklyReset,
	); f != nil {
		failures = append(failures, f)
	}

	if result.FableSeen {
		if g := s.RecordFable(
			result.FablePercent,
			"",
			fableReset(result.FableReset),
		); g != nil {
			failures = append(failures, g)
		}
	}

	s.notifier.Notify()

	return failures
}
