package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func NewService(t *testing.T) *service.Service {
	t.Helper()
	result := ForInventory(t, inventory.New())
	result.UseIndex(IndexDirectory())

	return result
}
