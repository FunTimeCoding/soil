package runner

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func command(
	inventoryDirectory string,
	playbook string,
	hosts []string,
) []string {
	result := []string{constant.Command}

	if inventoryDirectory != "" {
		result = append(
			result,
			constant.InventoryFlag,
			constant.StaticInventory,
			constant.InventoryFlag,
			filepath.Join(inventoryDirectory, constant.InventoryFile),
		)
	}

	if len(hosts) > 0 {
		result = append(result, constant.LimitFlag, join.Comma(hosts))
	}

	return append(result, playbook)
}
