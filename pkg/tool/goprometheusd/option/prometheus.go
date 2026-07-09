package option

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"

type Prometheus struct {
	Address   string
	Version   string
	Inventory *inventory.Inventory
}
