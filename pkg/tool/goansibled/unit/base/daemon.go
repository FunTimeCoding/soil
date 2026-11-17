package base

import "github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		s.daemon = s.build()
	}

	return s.daemon
}
