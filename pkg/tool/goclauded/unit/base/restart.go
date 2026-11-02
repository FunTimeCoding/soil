package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) Restart() {
	s.t.Helper()
	daemon.New(
		s.option,
		logger.New(s.t.Context()),
		s.Reporter,
		prometheus.NewRegistry(),
	).Close()
}
