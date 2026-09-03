package option

import "github.com/funtimecoding/soil/pkg/tool/goprometheusd/inventory"

type Prometheus struct {
	Address       string
	ServiceTokens []string
	Version   string
	Inventory *inventory.Inventory
}
