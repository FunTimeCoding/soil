package option

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/inventory"

type Alertmanager struct {
	Address   string
	Version   string
	Inventory *inventory.Inventory
}
