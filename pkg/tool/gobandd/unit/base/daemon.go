package base

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.Engine = func(i *instance.Instance) (face.BandClient, error) {
		return s.Instance(i.Name), nil
	}

	if s.Netbox != nil {
		o.Netbox = func() netbox.Client { return s.Netbox }
	}

	s.daemon = daemon.New(o, logger.New(s.t.Context()), s.Reporter)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
