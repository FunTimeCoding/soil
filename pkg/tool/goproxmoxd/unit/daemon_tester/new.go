package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"testing"
)

func New(
	t *testing.T,
	clients map[string]*mock_client.Client,
) *daemon.Daemon {
	t.Helper()

	return build(t, newOption(clients))
}
