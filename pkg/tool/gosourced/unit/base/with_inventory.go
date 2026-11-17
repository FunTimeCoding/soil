package base

import "github.com/funtimecoding/soil/pkg/source/inventory"

func WithInventory(i *inventory.Inventory) Option {
	return func(s *Server) {
		s.inventory = i
	}
}
