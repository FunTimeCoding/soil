package instance_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Inventory(names ...string) *inventory.Inventory[Instance, *Instance] {
	var instances []Instance

	for _, n := range names {
		instances = append(instances, *New(n, join.Empty(n, ".example")))
	}

	return inventory.New[Instance](instances...)
}
