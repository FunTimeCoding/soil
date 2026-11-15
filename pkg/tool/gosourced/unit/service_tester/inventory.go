package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
)

func Inventory(directory string) *inventory.Inventory {
	result := inventory.New()
	result.Modules = []module.Module{{Name: "echo", Directory: directory}}

	return result
}
