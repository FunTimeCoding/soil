package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"maps"
	"slices"
)

func (o *Tester) Inventory() *instance.Inventory {
	var result []instance.Instance

	for index, name := range slices.Sorted(maps.Keys(o.instances)) {
		result = append(
			result,
			instance.Instance{
				Name:  name,
				Host:  join.Empty(name, ".example"),
				Index: index,
			},
		)
	}

	return inventory.New[instance.Instance](result...)
}
