package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Workspace = s.t.TempDir()
	o.Stream = s.stream
	s.daemon = daemon.New(o, logger.New(s.t.Context()), memory.New())
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
