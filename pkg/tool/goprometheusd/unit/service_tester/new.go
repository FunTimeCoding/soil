package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"maps"
	"slices"
	"testing"
	"time"
)

func New(
	t *testing.T,
	clients map[string]*mock_client.Client,
) *service.Service {
	t.Helper()
	var instances []instance.Instance

	for _, name := range slices.Sorted(maps.Keys(clients)) {
		instances = append(
			instances,
			instance.Instance{Name: name, Host: join.Empty(name, ".example")},
		)
	}

	l := logger.New(t.Context())
	r := memory.New()
	p := pool.New(
		inventory.New[instance.Instance](instances...),
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

	return service.New(p, time.Now)
}
