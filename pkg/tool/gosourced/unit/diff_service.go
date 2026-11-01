package unit

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func diffService(directory string) *service.Service {
	i := inventory.New()
	i.Modules = []module.Module{{Name: "echo", Directory: directory}}

	return service.New(i)
}
