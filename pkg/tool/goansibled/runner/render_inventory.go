package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"path/filepath"
)

func (r *Runner) renderInventory() (string, error) {
	if r.inventory == nil {
		return "", nil
	}

	markup, e := r.inventory.AnsibleInventoryMarkup(
		context.Background(),
		gonetboxdConstant.AnsibleField,
	)

	if e != nil {
		return "", e
	}

	result := system.TemporaryDirectory(constant.InventoryPattern)
	system.SaveFile(filepath.Join(result, constant.InventoryFile), markup)

	return result, nil
}
