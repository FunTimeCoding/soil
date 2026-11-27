package bidi_server

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"

func (s *Server) Last() *message.Request {
	s.t.Helper()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.received) == 0 {
		s.t.Fatal("no command received")
	}

	return s.received[len(s.received)-1]
}
