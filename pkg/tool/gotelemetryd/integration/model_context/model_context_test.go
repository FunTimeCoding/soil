package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
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

func TestModelContextSummarisesOutcomesForOneActor(t *testing.T) {
	o := model_context_tester.New(t)
	c := o.Client
	o.SeedOutcome(
		"deploy_check",
		"Blair",
		constant.OutcomeSuccess,
		`{"step":"alfa"}`,
	)
	o.SeedOutcome(
		"deploy_check",
		"Blair",
		constant.OutcomeSuccess,
		`{"step":"alfa"}`,
	)
	o.SeedOutcome("deploy_check", "Blair", constant.OutcomeError, `{"step":"bravo"}`)
	o.SeedOutcome("deploy_check", "Cedar", constant.OutcomeError, `{"step":"bravo"}`)
	o.SeedOutcome("fleet_deploy", "Blair", constant.OutcomeError, `{"step":"bravo"}`)
	summary := c.MustCallTool(
		constant.Summary,
		map[string]any{
			constant.GroupBy: constant.Outcome,
			constant.Tool:    "deploy_check",
			constant.Actor:   "Blair",
		},
	)
	assert.StringContains(
		t,
		"\"count\": 2,\n\t\t\"outcome\": \"success\"",
		summary,
	)
	assert.StringContains(
		t,
		"\"count\": 1,\n\t\t\"outcome\": \"error\"",
		summary,
	)
	assert.StringNotContains(t, "fleet_deploy", summary)
	query := c.MustCallTool(
		constant.Query,
		map[string]any{constant.Tool: "deploy_check", constant.Actor: "Cedar"},
	)
	assert.StringContains(t, "\"detail\": {\n\t\t\t\"step\": \"bravo\"", query)
}
