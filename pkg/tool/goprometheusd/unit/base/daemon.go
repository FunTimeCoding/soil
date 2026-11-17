package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.Clock = s.Clock
	o.Client = func(i *instance.Instance) (face.PrometheusClient, error) {
		return s.Instance(i.Name), nil
	}
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
