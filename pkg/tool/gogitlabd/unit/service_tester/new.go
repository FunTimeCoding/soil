package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"testing"
)

func New(
	t *testing.T,
	names ...string,
) (*service.Service, map[string]*mock_client.Client) {
	t.Helper()
	clients := make(map[string]*mock_client.Client)
	var instances []instance.Instance

	for _, name := range names {
		clients[name] = mock_client.New()
		instances = append(
			instances,
			instance.Instance{Name: name, Host: join.Empty(name, ".example")},
		)
	}

	l := logger.New(t.Context())
	r := memory.New()
	i := inventory.New[instance.Instance](instances...)
	p := pool.New(
		i,
		func(i *instance.Instance) (*peered.Peered, error) {
			c := clients[i.Name]

			return peered.New(
				c,
				peer.New(
					join.Colon(constant.InstancePeer, i.Name),
					i.Host,
					c,
					l,
					r,
				),
			), nil
		},
	)
	t.Cleanup(p.Close)

	return service.New(p, i), clients
}
