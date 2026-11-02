package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func IntentService(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
	machines ...*client.VirtualMachine,
) (*service.Service, *netbox.Client) {
	t.Helper()
	n := netbox.New()
	t.Cleanup(n.Close)
	n.VirtualMachineList = append(n.VirtualMachineList, machines...)
	g := peer.New(
		gonetboxdConstant.Identity.Name(),
		n.Address(),
		n,
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(g.Close)
	h := sshMock.New()
	t.Cleanup(h.Close)

	return newService(t, c, h, r).WithNetbox(n, g), n
}
