package base

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Source = func() face.ItermClient { return s.Source }
	o.KeyInterval = 0
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
