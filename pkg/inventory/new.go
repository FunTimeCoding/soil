package inventory

import "github.com/funtimecoding/soil/pkg/inventory/face"

func New[T any, P face.Entry[T]](instances ...T) *Inventory[T, P] {
	return &Inventory[T, P]{Instances: instances}
}
