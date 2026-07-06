package service

import "github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"

func New(i *inventory.Inventory) *Service {
	return &Service{inventory: i}
}
