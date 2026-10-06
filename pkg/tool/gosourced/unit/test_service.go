package unit

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func testService() *service.Service {
	result := service.New(inventory.New())
	result.UseIndex(indexDirectory())

	return result
}
