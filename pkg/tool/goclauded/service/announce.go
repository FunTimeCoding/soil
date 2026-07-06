package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) Announce(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) error {
	if e := s.store.Announce(name, topic, files); e != nil {
		return e
	}

	if e := s.store.DeletePendingQueue(
		name,
		constant.QueueReannounce,
	); e != nil {
		return e
	}

	if e := s.store.DeletePendingQueue(
		name,
		constant.QueueTimeout,
	); e != nil {
		return e
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Announce,
		name,
		map[string]string{"topic": topic},
	); e != nil {
		return e
	}

	return s.PushQueueBroadcast(
		constant.QueueSessionAnnounce,
		fmt.Sprintf("%s - %s", name, topic),
	)
}
