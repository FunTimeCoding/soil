package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
	d constant.Delivery,
) (bool, error) {
	if fromName == "" {
		if f := entryRefusal(constant.QueuePulse, body); f != nil {
			return false, f
		}
	}

	if e := s.store.SendPulse(sessionIdentifier, fromName, body); e != nil {
		return false, e
	}

	if fromName == "" {
		callsign, e := s.store.CallsignBySessionIdentifier(sessionIdentifier)

		if e != nil {
			return false, e
		}

		if callsign != "" {
			return s.pushDelivery(
				queue.New(
					sessionIdentifier,
					callsign,
					constant.QueuePulse,
					body,
				),
				d,
			)
		}
	}

	s.notify()

	return false, nil
}
