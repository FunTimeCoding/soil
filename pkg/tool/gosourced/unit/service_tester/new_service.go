package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"testing"
)

func NewService(t *testing.T) *Services {
	t.Helper()
	result := InventoryServices(t, inventory.New())
	result.Workspace.UseIndex(IndexDirectory())

	return result
}
