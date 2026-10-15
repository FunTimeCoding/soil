package runner

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func (r *Runner) renderInventory() (string, error) {
	if r.inventory == nil {
		return "", nil
	}

	markup, e := r.inventory.Markup()

	if e != nil {
		return "", e
	}

	result := system.TemporaryDirectory(constant.InventoryPattern)
	system.SaveFile(filepath.Join(result, constant.InventoryFile), markup)

	return result, nil
}
