package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	clients   map[string]*alertmanager.Client
	sessions  sync.Map
	mutex     sync.Mutex
}
