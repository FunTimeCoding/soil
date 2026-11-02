package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/ssh/face"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"testing"
)

func NewWithShells(
	t *testing.T,
	clients map[string]*mock_client.Client,
	shells map[string]*sshMock.Client,
) *daemon.Daemon {
	t.Helper()
	o := newOption(clients)
	o.Shell = func(i *instance.Instance) (face.SSHClient, error) {
		return shells[i.Name], nil
	}

	return build(t, o)
}
