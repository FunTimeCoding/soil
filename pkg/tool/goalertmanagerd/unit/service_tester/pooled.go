package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/peered"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"testing"
)

func pooled(
	t *testing.T,
	i *instance.Inventory,
	alertmanager func(*instance.Instance) face.AlertmanagerClient,
	prometheus func(*instance.Instance) face.PrometheusClient,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	p := pool.New(
		i,
		func(i *instance.Instance) (*peered.Peered, error) {
			a := alertmanager(i)
			b := prometheus(i)

			return peered.New(
				a,
				b,
				peer.New(
					join.Colon(constant.AlertmanagerPeer, i.Name),
					i.Alertmanager.Host,
					a,
					l,
					r,
				),
				peer.New(
					join.Colon(constant.PrometheusPeer, i.Name),
					i.Prometheus.Host,
					b,
					l,
					r,
				),
			), nil
		},
	)
	t.Cleanup(p.Close)

	return service.New(p)
}
