package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := mock_client.New()
	c.Failure = outage_tester.RefusedDial()
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example"},
	)
	o.Client = func(*instance.Instance) (face.PrometheusClient, error) {
		return c, nil
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
		join.Colon(constant.InstancePeer, "alfa"),
	).Unavailable
}
