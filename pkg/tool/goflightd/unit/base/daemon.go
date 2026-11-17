package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.LitePath = s.litePath
	o.Command = s.System.Run
	o.Stream = s.System.Stream
	o.Clock = s.clock
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
