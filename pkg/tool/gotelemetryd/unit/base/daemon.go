package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.LitePath = s.litePath
	o.Clock = s.clock
	s.daemon = daemon.New(o, logger.New(context.Background()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
