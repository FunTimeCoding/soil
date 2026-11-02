package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
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
	l := logger.New(t.Context())
	p := pool.New(
		i,
		func(n *instance.Instance) (*peered.Peered, error) {
			return peered.New(
				c,
				peer.New(
					join.Colon(constant.InstancePeer, n.Name),
					n.Host,
					c,
					l,
					r,
				),
			), nil
		},
	)
	t.Cleanup(p.Close)
	t.Cleanup(shells.Close)
	s := pool.New(i, shells.build)
	t.Cleanup(s.Close)

	return service.New(p, s, i, l, r)
}
