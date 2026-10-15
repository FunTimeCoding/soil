package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func command(
	inventoryDirectory string,
	playbook string,
) []string {
	if inventoryDirectory == "" {
		return []string{constant.Command, playbook}
	}

	return []string{
		constant.Command,
		constant.InventoryFlag,
		constant.StaticInventory,
		constant.InventoryFlag,
		filepath.Join(inventoryDirectory, constant.InventoryFile),
		playbook,
	}
}
