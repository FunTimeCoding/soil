package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func clockedInstance(
	t *testing.T,
) (*model_context_client.Client, *mock_client.Client) {
	t.Helper()
	m := mock_client.New()
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			daemon_tester.NewWith(
				t,
				map[string]*mock_client.Client{
					"alfa":  m,
					"bravo": mock_client.New(),
				},
				windowNow,
			).Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(v.Close)
	c := v.Connect(t)
	t.Cleanup(c.Close)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)

	return c, m
}
