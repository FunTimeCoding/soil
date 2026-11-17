package base

import (
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/daemon"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Source = func() face.LinkAceClient { return s.Source }
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
