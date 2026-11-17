package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogated/daemon"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	s.daemon = daemon.New(
		s.daemonOption(),
		logger.New(context.Background()),
		s.Reporter,
	)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
