package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (s *notifySink) all() []client.NotifyRequest {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return append([]client.NotifyRequest{}, s.received...)
}
