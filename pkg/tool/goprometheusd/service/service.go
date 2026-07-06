package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	clients   map[string]*prometheus.Client
	sessions  sync.Map
	mutex     sync.Mutex
}
