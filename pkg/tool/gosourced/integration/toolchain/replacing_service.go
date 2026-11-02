package toolchain

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func replacingService(
	t *testing.T,
	library string,
	user string,
) *service.Service {
	t.Helper()
	i := inventory.New()
	i.Add("lib", library)
	i.Add("user", user)
	result := service_tester.ForInventory(t, i)
	result.UseIndex(service_tester.IndexDirectory())

	return result
}
