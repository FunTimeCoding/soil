package roster

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time/minute"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) sweepInactivity() []error {
	var result []error
	cutoff := s.clock().Add(-constant.InactivityTimeoutWindow)
	sessions := s.store.SweepInactivity(cutoff)

	for _, e := range sessions {
		if f := s.store.LogEvent(
			e.Identifier,
			constant.InactivityTimeout,
			e.CallsignValue(),
			map[string]string{"topic": e.Topic},
		); f != nil {
			result = append(result, f)
		}

		if f := s.message.PushQueue(
			e.Identifier,
			e.CallsignValue(),
			constant.QueueTimeout,
			fmt.Sprintf(
				"%s since last turn. Removed from roster. Last topic: %s",
				minute.Readable(int(constant.InactivityTimeoutWindow.Minutes())),
				e.Topic,
			),
		); f != nil {
			result = append(result, f)
		}
	}

	if len(sessions) > 0 {
		s.notify()
	}

	return result
}
