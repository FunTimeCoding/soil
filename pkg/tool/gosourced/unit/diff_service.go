package unit

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func diffService(
	t *testing.T,
	directory string,
) *service.Service {
	t.Helper()
	i := inventory.New()
	i.Modules = []module.Module{{Name: "echo", Directory: directory}}

	return service_tester.ForInventory(t, i)
}
