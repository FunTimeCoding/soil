package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func adHoc(
	inventoryDirectory string,
	pattern string,
	arguments ...string,
) []string {
	result := []string{constant.HostCommandProgram}

	if inventoryDirectory != "" {
		result = append(
			result,
			constant.InventoryFlag,
			constant.StaticInventory,
			constant.InventoryFlag,
			filepath.Join(inventoryDirectory, constant.InventoryFile),
		)
	}

	return append(append(result, pattern), arguments...)
}
