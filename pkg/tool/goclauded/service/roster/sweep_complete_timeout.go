package roster

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) sweepCompleteTimeout() []error {
	var result []error
	cutoff := s.clock().Add(-constant.CompleteTimeoutWindow)
	sessions := s.store.SweepCompleteTimeout(cutoff)

	for _, e := range sessions {
		if f := s.store.LogEvent(
			e.Identifier,
			constant.CompleteTimeout,
			e.CallsignValue(),
			nil,
		); f != nil {
			result = append(result, f)
		}

		if f := s.message.PushQueue(
			e.Identifier,
			e.CallsignValue(),
			constant.QueueTimeout,
			"30 minutes since completing. Removed from roster.",
		); f != nil {
			result = append(result, f)
		}
	}

	if len(sessions) > 0 {
		s.notify()
	}

	return result
}
