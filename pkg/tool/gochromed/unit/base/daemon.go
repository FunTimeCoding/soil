package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		o := option.New()
		o.DownloadDirectory = s.t.TempDir()
		o.Source = func() face.ChromiumSource { return s.Source }
		s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
