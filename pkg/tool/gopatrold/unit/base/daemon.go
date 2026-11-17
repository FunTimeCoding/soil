package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Instances = s.names
	o.Source = func(instance string) face.PatrolClient {
		return s.instances[instance]
	}
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter, s.Metric)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
