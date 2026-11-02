package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
)

func New(
	p *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered],
	i *instance.Inventory,
) *Service {
	return &Service{
		inventory: i,
		pool:      p,
		notifier:  notifier.New(),
		entries:   make(map[string][]*board_entry.Entry),
		requests:  make(map[string][]*request_entry.Entry),
	}
}
