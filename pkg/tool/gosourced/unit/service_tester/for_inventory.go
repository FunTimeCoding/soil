package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"testing"
)

func ForInventory(
	t *testing.T,
	i *inventory.Inventory,
) *Services {
	t.Helper()
	m := mock_client.New()
	t.Cleanup(m.Close)

	return fromParts(t, i, m)
}
