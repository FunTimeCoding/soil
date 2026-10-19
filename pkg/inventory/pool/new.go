package pool

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/face"
)

func New[T any, P face.Entry[T], C any](
	i *inventory.Inventory[T, P],
	build func(*T) (C, error),
) *Pool[T, P, C] {
	return &Pool[T, P, C]{
		inventory: i,
		build:     build,
		clients:   make(map[string]C),
		built:     make(map[string]bool),
	}
}
