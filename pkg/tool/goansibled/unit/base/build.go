package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) build() *daemon.Daemon {
	s.t.Helper()
	s.Registry = prometheus.NewRegistry()
	result := daemon.New(
		s.Option,
		logger.New(s.t.Context()),
		s.Reporter,
		s.Registry,
	)
	s.t.Cleanup(
		func() {
			s.Release()
			result.Close()
		},
	)

	return result
}
