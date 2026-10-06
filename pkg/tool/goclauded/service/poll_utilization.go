package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"

func (s *Service) PollUtilization() {
	c := utilization.ReadCredential()

	if c == nil {
		s.logger.Structured("utilization poll skipped", "reason", "credential")

		return
	}

	now := s.clock()

	if c.Expired(now) {
		s.logger.Structured(
			"utilization poll skipped",
			"reason",
			"expired",
			"expiredAt",
			c.ExpiresAt,
		)

		return
	}

	result, e := utilization.Read(c.AccessToken)

	if e != nil {
		s.logger.Structured("utilization poll failed", "error", e)

		return
	}

	if e := s.RecordRateLimits(
		result.SessionPercent,
		result.WeeklyPercent,
		result.SessionReset,
		result.WeeklyReset,
	); e != nil {
		s.logger.Structured("rate snapshot failed", "error", e)
	}

	if result.FableSeen {
		if e := s.RecordFable(
			result.FablePercent,
			"",
			fableReset(result.FableReset),
		); e != nil {
			s.logger.Structured("fable snapshot failed", "error", e)
		}
	}

	s.notifier.Notify()
}
