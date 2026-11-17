package base

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"time"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	n := notifier.New()
	o := option.New()
	o.LitePath = s.litePath
	o.Clock = time.Now
	o.Notifier = func() face.EventNotifier { return n }
	o.SeedDirectory = s.seedDirectory
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
