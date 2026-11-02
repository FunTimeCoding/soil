package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func ContainerService(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	h := sshMock.New()
	t.Cleanup(h.Close)

	return newService(t, c, h, memory.New())
}
