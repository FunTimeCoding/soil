package option

import "github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"

type Source struct {
	Address   string
	Version   string
	Inventory *inventory.Inventory
}
