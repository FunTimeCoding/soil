package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/ssh/face"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func NewPollWorker(
	t *testing.T,
	name string,
	c *mock_client.Client,
) (*worker.Worker, *prometheus.Registry, *memory.Memory) {
	t.Helper()
	y := prometheus.NewRegistry()
	r := memory.New()
	l := logger.New(t.Context())
	i := inventory.New[instance.Instance](
		instance.Instance{Name: name, Host: "mock"},
	)
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
	shells := pool.New(
		i,
		func(*instance.Instance) (face.SSHClient, error) {
			return sshMock.New(), nil
		},
	)
	t.Cleanup(shells.Close)
	v := service.New(p, shells, i, l, r)
	result := worker.New(v.Refresh, time.Minute, y, l, r)
	t.Cleanup(result.Close)

	return result, y, r
}
