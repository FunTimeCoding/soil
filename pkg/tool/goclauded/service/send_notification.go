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
) error {
	holder, e := s.store.SessionByCallsign(callsign)

	if e != nil {
		return e
	}

	if holder == nil {
		return not_found.New(constant.Callsign, callsign)
	}

	if e := s.store.SendNotification(
		holder.Identifier,
		callsign,
		source,
		body,
	); e != nil {
		return e
	}

	return s.PushQueue(
		holder.Identifier,
		callsign,
		constant.QueueNotification,
		fmt.Sprintf("%s: %s", source, body),
	)
}
