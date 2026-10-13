package inventory

import "github.com/funtimecoding/soil/pkg/inventory/face"

type Inventory[T any, P face.Entry[T]] struct {
	Instances []T `yaml:"instances"`
}
