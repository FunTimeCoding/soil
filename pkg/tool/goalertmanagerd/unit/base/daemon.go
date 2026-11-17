package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.Alertmanager = func(i *instance.Instance) (face.AlertmanagerClient, error) {
		return s.Instance(i.Name).Alertmanager, nil
	}
	o.Prometheus = func(i *instance.Instance) (face.PrometheusClient, error) {
		return s.Instance(i.Name).Prometheus, nil
	}
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
