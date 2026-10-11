package toolchain

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
)

func replacingService(
	library string,
	user string,
) *service.Service {
	i := inventory.New()
	i.Add("lib", library)
	i.Add("user", user)
	result := service.New(i)
	result.UseIndex(service_tester.IndexDirectory())

	return result
}
