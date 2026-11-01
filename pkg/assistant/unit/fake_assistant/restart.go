package fake_assistant

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Server) Restart() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.socket == nil {
		return
	}

	errors.LogClose(s.socket)
	s.socket = nil
}
