package service_tester

import "github.com/funtimecoding/soil/pkg/source/inventory"

func WithInventory(i *inventory.Inventory) Option {
	return func(s *setup) {
		s.inventory = i
	}
}
