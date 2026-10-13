package option

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"

type Postgres struct {
	Address       string
	ServiceTokens []string
	Inventory     *instance.Inventory
}
