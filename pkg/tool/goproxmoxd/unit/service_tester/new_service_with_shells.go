package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"testing"
)

func newServiceWithShells(
	t *testing.T,
	c *mock_client.Client,
	shells *Shells,
	r face.Reporter,
) *service.Service {
	t.Helper()
	i := containerOption().Inventory
	p := pool.New(
		i,
		func(*instance.Instance) (proxmox.ProxmoxClient, error) {
			return c, nil
		},
	)
	t.Cleanup(p.Close)
	t.Cleanup(shells.Close)
	s := pool.New(i, shells.build)
	t.Cleanup(s.Close)

	return service.New(p, s, i, logger.New(t.Context()), r)
}
