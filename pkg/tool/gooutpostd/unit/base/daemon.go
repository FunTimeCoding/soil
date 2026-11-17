package base

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/daemon"

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		s.daemon = daemon.New()
	}

	return s.daemon
}
