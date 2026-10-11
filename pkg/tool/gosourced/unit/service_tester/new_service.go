package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func NewService() *service.Service {
	result := service.New(inventory.New())
	result.UseIndex(IndexDirectory())

	return result
}
