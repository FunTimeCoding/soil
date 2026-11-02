package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, proxmox.ProxmoxClient],
	shells *pool.Pool[instance.Instance, *instance.Instance, ssh.SSHClient],
	i *instance.Inventory,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	return &Service{
		inventory: i,
		pool:      p,
		shells:    shells,
		logger:    l,
		reporter:  r,
	}
}
