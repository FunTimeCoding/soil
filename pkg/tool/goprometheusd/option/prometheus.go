package option

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"

type Prometheus struct {
	Port      int
	Version   string
	Inventory *inventory.Inventory
}
