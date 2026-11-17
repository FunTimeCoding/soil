package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.BackupDirectory = s.BackupDirectory
	o.Dump = mock_dump.New()
	o.Client = func(i *instance.Instance) (face.PostgresClient, error) {
		return s.Instance(i.Name), nil
	}
	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
