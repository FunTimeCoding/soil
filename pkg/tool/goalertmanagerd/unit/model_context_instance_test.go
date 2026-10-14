package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestModelContextAlertsNeedASelectedInstance(t *testing.T) {
	alfa := instance.Instance{Name: "alfa"}
	alfa.Alertmanager.Host = "alertmanager.alfa.example"
	alfa.Prometheus.Host = "prometheus.alfa.example"
	bravo := instance.Instance{Name: "bravo"}
	bravo.Alertmanager.Host = "alertmanager.bravo.example"
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goalertmanagerd.Mount(
				service.New(inventory.New[instance.Instance](alfa, bravo)),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	c := v.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"no instance selected",
		c.MustCallToolError(constant.ListAlerts, nil),
	)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "alfa"},
	)
	var listed []instance_response.Response
	notation.MustDecodeBytes(
		[]byte(c.MustCallTool(inventoryConstant.ListInstances, nil)),
		&listed,
		false,
	)
	assert.Count(t, 2, listed)
	assert.True(t, listed[0].Active)
	assert.String(t, "alertmanager.alfa.example", listed[0].Host)
	assert.False(t, listed[1].Active)
}
