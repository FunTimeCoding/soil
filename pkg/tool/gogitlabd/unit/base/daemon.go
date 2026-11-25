package base

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.Client = func(i *instance.Instance) (face.GitLabClient, error) {
		return s.Instance(i.Name), nil
	}
	o.LitePath = s.Path

	if c := s.Connector(); c != nil {
		o.Connector = func() goclauded.Client { return c }
	}

	s.daemon = daemon.New(
		o,
		logger.New(s.t.Context()),
		s.Reporter,
		prometheus.NewRegistry(),
	)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
