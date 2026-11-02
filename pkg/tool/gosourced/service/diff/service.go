package diff

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module"
)

type Service struct {
	inventory *inventory.Inventory
	module    *module.Service
}
