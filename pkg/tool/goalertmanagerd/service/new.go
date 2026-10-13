package service

import (
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/inventory/selection"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func New(i *instance.Inventory) *Service {
	return &Service{
		inventory: i,
		pool:      pool.New(i, newClient),
		selection: selection.New(),
	}
}
