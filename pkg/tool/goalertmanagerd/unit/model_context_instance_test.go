package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
)

func TestModelContextAlertsNeedSelectedInstance(t *testing.T) {
	c := base.New(t, service_tester.WithInstances("alfa", "bravo")).Tools()
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
