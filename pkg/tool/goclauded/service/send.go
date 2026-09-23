package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) Send(
	name string,
	to string,
	body string,
	immediate bool,
) (bool, error) {
	var holder *session.Session

	if to != "" {
		found, e := s.store.SessionByCallsign(to)

		if e != nil {
			return false, e
		}

		if found == nil {
			return false, not_found.New(constant.Callsign, to)
		}

		holder = found
	}

	if e := s.store.SendMessage(name, to, body); e != nil {
		return false, e
	}

	formatted := fmt.Sprintf("%s: %s", name, body)

	if holder == nil {
		return false, s.PushQueueBroadcast(constant.QueueMessage, formatted)
	}

	if immediate {
		return s.PushQueueImmediate(
			holder.Identifier,
			to,
			constant.QueueMessage,
			formatted,
		)
	}

	return false, s.PushQueue(
		holder.Identifier,
		to,
		constant.QueueMessage,
		formatted,
	)
}
