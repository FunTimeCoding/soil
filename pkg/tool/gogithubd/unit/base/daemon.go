package base

import (
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Owner = "alfa"
	o.Source = func() face.GitHubClient { return s.Source }
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter, s.Metric)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
