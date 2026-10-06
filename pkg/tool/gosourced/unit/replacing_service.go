package unit

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func replacingService(
	library string,
	user string,
) *service.Service {
	i := inventory.New()
	i.Add("lib", library)
	i.Add("user", user)
	result := service.New(i)
	result.UseIndex(indexDirectory())

	return result
}
