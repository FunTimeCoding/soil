package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) SendNotification(
	callsign string,
	source string,
	body string,
) error {
	if e := s.store.SendNotification(callsign, source, body); e != nil {
		return e
	}

	return s.PushQueue(
		callsign,
		constant.QueueNotification,
		fmt.Sprintf("%s: %s", source, body),
	)
}
