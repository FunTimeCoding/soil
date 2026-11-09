package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Inventory() {
	i.String(
		constant.Inventory,
		environment.Fallback(
			inventoryConstant.PathEnvironment,
			i.identity.InventoryPath(),
		),
		inventoryConstant.PathUsage,
	)
}
