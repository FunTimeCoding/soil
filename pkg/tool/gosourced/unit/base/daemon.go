package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon == nil {
		o := option.New()
		o.Inventory = s.inventory
		o.IndexDirectory = s.indexDirectory
		o.Memory = func() face.Client { return s.Memory }
		s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
		s.t.Cleanup(s.daemon.Close)
	}

	return s.daemon
}
