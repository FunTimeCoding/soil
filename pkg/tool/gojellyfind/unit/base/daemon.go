package base

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Source = func() face.JellyfinClient { return s.source }
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
