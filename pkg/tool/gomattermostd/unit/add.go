package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (s *notifySink) add(v client.NotifyRequest) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.received = append(s.received, v)
}
