package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) PushQueueBroadcast(
	kind string,
	body string,
) error {
	return s.pushEntryBroadcast(queue.NewBroadcast(kind, body))
}
