package option

import "github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"

type Source struct {
	Port      int
	Version   string
	Inventory *inventory.Inventory
}
