package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	c := o.Client
	assert.Count(t, 2, c.ListTools())
	o.Seed("save_memory", constant.SurfaceModelContext, "Blair")
	o.Seed("save_memory", constant.SurfaceModelContext, "Cedar")
	o.Seed("fleet_deploy", constant.SurfaceCommandLine, "Blair")
	query := c.MustCallTool(constant.Query, map[string]any{})
	assert.StringContains(t, "save_memory", query)
	assert.StringContains(t, "fleet_deploy", query)
	filtered := c.MustCallTool(
		constant.Query,
		map[string]any{constant.Tool: "save_memory"},
	)
	assert.StringContains(t, "save_memory", filtered)
	assert.StringNotContains(t, "fleet_deploy", filtered)
	summary := c.MustCallTool(constant.Summary, map[string]any{})
	assert.StringContains(t, "save_memory", summary)
	assert.StringContains(t, "fleet_deploy", summary)
}
