package option

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory"

type Postgres struct {
	Address       string
	ServiceTokens []string
	Inventory *inventory.Inventory
	Version   string
}
