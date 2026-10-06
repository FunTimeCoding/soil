package option

import "github.com/funtimecoding/soil/pkg/source/inventory"

type Source struct {
	Address       string
	ServiceTokens []string
	Inventory     *inventory.Inventory
}
