package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
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
		notifier:  notifier.New(),
	}
}
