package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
	immediate bool,
) (bool, error) {
	if e := s.store.SendPulse(sessionIdentifier, fromName, body); e != nil {
		return false, e
	}

	if fromName == "" {
		callsign, e := s.store.CallsignBySessionIdentifier(sessionIdentifier)

		if e != nil {
			return false, e
		}

		if callsign != "" {
			if immediate {
				return s.PushQueueImmediate(
					sessionIdentifier,
					callsign,
					constant.QueuePulse,
					body,
				)
			}

			return false, s.PushQueue(
				sessionIdentifier,
				callsign,
				constant.QueuePulse,
				body,
			)
		}
	}

	s.notify()

	return false, nil
}
