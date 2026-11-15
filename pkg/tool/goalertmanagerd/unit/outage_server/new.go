package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	a := mock_client.New()
	t.Cleanup(a.Close)
	p := prometheusMock.New()
	t.Cleanup(p.Close)

	switch peer {
	case constant.AlertmanagerPeer:
		a.Fail(outage_tester.RefusedDial())
	case constant.PrometheusPeer:
		p.Failure = outage_tester.RefusedDial()
	}

	i := instance.Instance{Name: "alfa"}
	i.Alertmanager.Host = "alertmanager.alfa.example"
	i.Prometheus.Host = "prometheus.alfa.example"
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](i)
	o.Alertmanager = func(*instance.Instance) (face.AlertmanagerClient, error) {
		return a, nil
	}
	o.Prometheus = func(*instance.Instance) (face.PrometheusClient, error) {
		return p, nil
	}
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(
		t,
		d.Workers(),
		join.Colon(peer, i.Name),
	).Unavailable
}
