package base

import (
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Source = func() face.BackrestClient { return s.Source }
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter, s.Metric)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
