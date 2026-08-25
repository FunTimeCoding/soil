package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) Release(
	sessionIdentifier string,
	name string,
) error {
	if e := s.PushQueueBroadcast(
		constant.QueueSessionRelease,
		fmt.Sprintf("%s released", name),
	); e != nil {
		return e
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Release,
		name,
		nil,
	); e != nil {
		return e
	}

	return s.store.ReleaseByCallsign(name)
}
