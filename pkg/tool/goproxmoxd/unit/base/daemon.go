package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	o := option.New()
	o.Inventory = s.Inventory()
	o.Hypervisor = func(i *instance.Instance) (proxmox.ProxmoxClient, error) {
		return s.Instance(i.Name).Hypervisor, nil
	}
	o.Shell = func(i *instance.Instance) (ssh.SSHClient, error) {
		return s.Instance(i.Name).Shell, nil
	}

	if s.Netbox != nil {
		o.Netbox = func() netbox.Client { return s.Netbox }
	}

	s.daemon = daemon.New(
		o,
		logger.New(s.t.Context()),
		s.Reporter,
		prometheus.NewRegistry(),
	)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
