package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/daemon"
	"testing"
)

func New(
	t *testing.T,
	i *inventory.Inventory,
) (*daemon.Daemon, *mock_client.Client) {
	t.Helper()

	return NewInIndex(t, i, t.TempDir())
}
