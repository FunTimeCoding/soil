package pool

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/face"
	"sync"
)

type Pool[T any, P face.Entry[T], C any] struct {
	inventory *inventory.Inventory[T, P]
	build     func(*T) (C, error)
	clients   map[string]C
	built     map[string]bool
	mutex     sync.Mutex
}
