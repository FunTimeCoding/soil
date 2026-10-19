package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/unit/session_tester"
	"github.com/funtimecoding/soil/pkg/inventory"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/mark3labs/mcp-go/server"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
	"time"
)

func TestModelContextNodesFollowTheSelectedInstance(t *testing.T) {
	s := service.New(
		inventory.New[instance.Instance](
			instance.Instance{Name: "alfa", Host: "alfa.example"},
			instance.Instance{Name: "bravo", Host: "bravo.example", Index: 1},
		),
	)
	a := mock_client.New()
	a.AddNode("alfa-node")
	s.Pool().Set("alfa", a)
	b := mock_client.New()
	b.AddNode("bravo-node")
	s.Pool().Set("bravo", b)
	k := worker.New(
		s,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		memory.New(),
	)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goproxmoxd.Mount(
				s,
				web.New(s, k),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	c := v.Connect(t)
	defer c.Close()
	assert.StringContains(t, "use_instance", c.InitializeResult().Instructions)
	assert.StringContains(
		t,
		"no instance selected",
		c.MustCallToolError(constant.ListNodes, nil),
	)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "bravo"},
	)
	nodes := c.MustCallTool(constant.ListNodes, nil)
	assert.StringContains(t, "bravo-node", nodes)
	assert.StringNotContains(t, "alfa-node", nodes)
	var listed []instance_response.Response
	notation.MustDecodeBytes(
		[]byte(c.MustCallTool(inventoryConstant.ListInstances, nil)),
		&listed,
		false,
	)
	assert.Count(t, 2, listed)
	assert.False(t, listed[0].Active)
	assert.True(t, listed[1].Active)
}

func TestModelContextClosingForgetsTheSelection(t *testing.T) {
	s := service.New(
		inventory.New[instance.Instance](
			instance.Instance{Name: "alfa", Host: "alfa.example"},
			instance.Instance{Name: "bravo", Host: "bravo.example", Index: 1},
		),
	)
	k := worker.New(
		s,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		memory.New(),
	)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goproxmoxd.Mount(
				s,
				web.New(s, k),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	c := v.Connect(t)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "bravo"},
	)
	x := server.NewMCPServer("test", "0").WithContext(
		t.Context(),
		session_tester.New(c.SessionIdentifier()),
	)
	selected, e := s.Selection().Active(x)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", selected)
	c.Close()
	cleared, f := s.Selection().Active(x)
	assert.FatalOnError(t, f)
	assert.String(t, "", cleared)
}
