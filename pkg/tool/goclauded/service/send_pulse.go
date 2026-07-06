package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
) error {
	if e := s.store.SendPulse(sessionIdentifier, fromName, body); e != nil {
		return e
	}

	if fromName == "" {
		callsign, e := s.store.CallsignBySessionIdentifier(sessionIdentifier)

		if e != nil {
			return e
		}

		if callsign != "" {
			return s.PushQueue(callsign, constant.QueuePulse, body)
		}
	}

	s.notify()

	return nil
}
