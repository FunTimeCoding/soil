package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Source = func() face.OpenbaoClient { return s.Source }
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter, s.Registry)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
