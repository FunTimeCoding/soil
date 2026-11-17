package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/daemon"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		s.daemon = daemon.New(
			s.Option,
			logger.New(s.t.Context()),
			s.Reporter,
			s.Registry,
		)
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
