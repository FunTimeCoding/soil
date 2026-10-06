package option

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/inventory"

type Alertmanager struct {
	Address       string
	ServiceTokens []string
	Inventory     *inventory.Inventory
}
