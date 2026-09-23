package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func testService() *service.Service {
	return service.New(inventory.New())
}
