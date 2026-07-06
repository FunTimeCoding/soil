package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepCompleteTimeout() {
	cutoff := s.clock().Add(-30 * time.Minute)
	sessions := s.store.SweepCompleteTimeout(cutoff)

	for _, e := range sessions {
		errors.PanicOnError(
			s.store.LogEvent(
				e.Identifier,
				constant.CompleteTimeout,
				e.CallsignValue(),
				nil,
			),
		)
		errors.PanicOnError(
			s.PushQueue(
				e.CallsignValue(),
				constant.QueueTimeout,
				"30 minutes since completing. Removed from roster.",
			),
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
