package roster

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepInactivity() []error {
	var result []error
	cutoff := s.clock().Add(-1 * time.Hour)
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
				"1 hour since last turn. Removed from roster. Last topic: %s",
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
