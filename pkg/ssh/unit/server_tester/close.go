package server_tester

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Server) Close() {
	s.closed.Do(
		func() {
			errors.LogReclose(s.listener)
			s.mutex.Lock()

			for _, c := range s.connections {
				errors.LogReclose(c)
			}

			s.mutex.Unlock()
		},
	)
	s.group.Wait()
}
