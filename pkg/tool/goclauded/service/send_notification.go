package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) SendNotification(
	callsign string,
	source string,
	body string,
	immediate bool,
) (bool, error) {
	holder, e := s.store.SessionByCallsign(callsign)

	if e != nil {
		return false, e
	}

	if holder == nil {
		return false, not_found.New(constant.Callsign, callsign)
	}

	if e := s.store.SendNotification(
		holder.Identifier,
		callsign,
		source,
		body,
	); e != nil {
		return false, e
	}

	formatted := fmt.Sprintf("%s: %s", source, body)

	if immediate {
		return s.PushQueueImmediate(
			holder.Identifier,
			callsign,
			constant.QueueNotification,
			formatted,
		)
	}

	return false, s.PushQueue(
		holder.Identifier,
		callsign,
		constant.QueueNotification,
		formatted,
	)
}
