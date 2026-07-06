package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepInactivity() {
	cutoff := s.clock().Add(-1 * time.Hour)
	sessions := s.store.SweepInactivity(cutoff)

	for _, e := range sessions {
		errors.PanicOnError(
			s.store.LogEvent(
				e.Identifier,
				constant.InactivityTimeout,
				e.CallsignValue(),
				map[string]string{"topic": e.Topic},
			),
		)
		errors.PanicOnError(
			s.PushQueue(
				e.CallsignValue(),
				constant.QueueTimeout,
				fmt.Sprintf(
					"1 hour since last turn. Removed from roster. Last topic: %s",
					e.Topic,
				),
			),
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
