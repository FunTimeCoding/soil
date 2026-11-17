package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestModelContextNodesFollowSelectedInstance(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	s.Alfa().Hypervisor.AddNode("alfa-node")
	s.Bravo().Hypervisor.AddNode("bravo-node")
	c := s.Tools()
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
