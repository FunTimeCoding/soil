package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"maps"
	"slices"
)

func (o *Tester) Inventory() *instance.Inventory {
	var result []instance.Instance

	for _, name := range slices.Sorted(maps.Keys(o.instances)) {
		i := instance.Instance{Name: name}
		i.Alertmanager.Host = join.Empty("alertmanager.", name, ".example")
		i.Prometheus.Host = join.Empty("prometheus.", name, ".example")
		result = append(result, i)
	}

	return inventory.New[instance.Instance](result...)
}
