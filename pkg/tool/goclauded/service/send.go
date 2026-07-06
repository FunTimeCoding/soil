package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) Send(
	name string,
	to string,
	body string,
) error {
	if to != "" {
		session, e := s.store.SessionByCallsign(to)

		if e != nil {
			return e
		}

		if session == nil {
			return fmt.Errorf(
				"%w: %s",
				constant.ErrorCallsignNotFound,
				to,
			)
		}
	}

	if e := s.store.SendMessage(name, to, body); e != nil {
		return e
	}

	formatted := fmt.Sprintf("%s: %s", name, body)

	if to != "" {
		return s.PushQueue(to, constant.QueueMessage, formatted)
	}

	return s.PushQueueBroadcast(constant.QueueMessage, formatted)
}
