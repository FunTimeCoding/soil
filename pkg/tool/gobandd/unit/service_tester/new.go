package service_tester

import (
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"testing"
)

func New(
	t *testing.T,
	r face.Reporter,
	names ...string,
) (*service.Service, map[string]*mock_client.Client) {
	t.Helper()
	engines := make(map[string]*mock_client.Client)
	var instances []instance.Instance

	for _, name := range names {
		engines[name] = mock_client.New()
		instances = append(
			instances,
			instance.Instance{Name: name, Host: join.Empty(name, ".example")},
		)
	}

	l := logger.New(t.Context())
	p := pool.New(
		inventory.New[instance.Instance](instances...),
		func(i *instance.Instance) (*peered.Peered, error) {
			c := engines[i.Name]

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

	return service.New(p, l, r), engines
}
