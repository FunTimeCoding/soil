package option

import "github.com/funtimecoding/soil/pkg/source/inventory"

func (c *Check) ModuleInventory() *inventory.Inventory {
	if c.Inventory == nil {
		return inventory.LoadOptional(inventory.DefaultPath())
	}

	return c.Inventory
}
