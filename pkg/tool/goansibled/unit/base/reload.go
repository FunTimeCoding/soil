package base

import "github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"

func (s *Server) Reload() *daemon.Daemon {
	s.t.Helper()
	s.daemon = s.build()

	return s.daemon
}
