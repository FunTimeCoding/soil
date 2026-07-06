package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory: i,
		clients:   make(map[string]*prometheus.Client),
	}
}
