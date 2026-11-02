package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func newService(
	t *testing.T,
	c *mock_client.Client,
	h *sshMock.Client,
	r face.Reporter,
) *service.Service {
	t.Helper()

	return newServiceWithShells(t, c, newShells(h), r)
}
