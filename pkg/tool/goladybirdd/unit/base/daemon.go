package base

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		o := option.New()
		o.DownloadDirectory = s.t.TempDir()
		o.Client = func(events func(*message.Reply)) face.LadybirdClient {
			s.Client.Listen(events)

			return s.Client
		}
		s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
