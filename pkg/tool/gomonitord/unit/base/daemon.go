package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"path/filepath"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		o := option.New()
		o.LitePath = filepath.Join(s.t.TempDir(), "gomonitord.db")
		s.daemon = daemon.New(o, logger.New(s.t.Context()))
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
