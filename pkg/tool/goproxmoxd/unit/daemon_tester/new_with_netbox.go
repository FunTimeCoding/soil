package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"testing"
)

func NewWithNetbox(
	t *testing.T,
	clients map[string]*mock_client.Client,
	n *netbox.Client,
) *daemon.Daemon {
	t.Helper()
	o := newOption(clients)
	o.Netbox = func() face.Client { return n }

	return build(t, o)
}
