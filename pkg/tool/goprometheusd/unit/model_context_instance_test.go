package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/service_tester"
	"testing"
)

func TestModelContextQueriesNeedSelectedInstance(t *testing.T) {
	c := base.New(t, service_tester.WithInstances("alfa", "bravo")).Tools()
	assert.StringContains(
		t,
		"no instance selected",
		c.MustCallToolError(
			constant.Query,
			map[string]any{generativeConstant.ParameterQuery: "up"},
		),
	)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "bravo"},
	)
	var listed []instance_response.Response
	notation.MustDecodeBytes(
		[]byte(c.MustCallTool(inventoryConstant.ListInstances, nil)),
		&listed,
		false,
	)
	assert.Count(t, 2, listed)
	assert.False(t, listed[0].Active)
	assert.True(t, listed[1].Active)
	assert.String(t, "bravo.example", listed[1].Host)
}
