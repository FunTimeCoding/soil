package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestModelContextNodesFollowSelectedInstance(t *testing.T) {
	a := mock_client.New()
	defer a.Close()
	a.AddNode("alfa-node")
	b := mock_client.New()
	defer b.Close()
	b.AddNode("bravo-node")
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": a, "bravo": b},
	)
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
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
