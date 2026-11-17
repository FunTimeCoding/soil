package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		s.daemon = daemon.New(
			s.Option,
			logger.New(s.t.Context()),
			s.Reporter,
			prometheus.NewRegistry(),
		)
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
