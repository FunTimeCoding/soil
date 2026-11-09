package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"testing"
	"time"
)

func TestQueryMomentRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(
		constant.Query,
		map[string]any{generativeConstant.ParameterQuery: "up"},
	)
	moments := m.Moments()
	assert.Count(t, 1, moments)
	assert.String(
		t,
		"2026-08-16T20:00:00Z",
		moments[0].UTC().Format(time.RFC3339),
	)
}

func TestLabelNamesStartRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(constant.LabelNames, map[string]any{})
	moments := m.Moments()
	assert.Count(t, 2, moments)
	assert.String(
		t,
		"2026-08-16T19:00:00Z",
		moments[0].UTC().Format(time.RFC3339),
	)
}

func TestLabelNamesEndRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(constant.LabelNames, map[string]any{})
	moments := m.Moments()
	assert.Count(t, 2, moments)
	assert.String(
		t,
		"2026-08-16T20:00:00Z",
		moments[1].UTC().Format(time.RFC3339),
	)
}

func TestLabelValuesStartRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(constant.LabelValues, map[string]any{"label": "job"})
	moments := m.Moments()
	assert.Count(t, 2, moments)
	assert.String(
		t,
		"2026-08-16T19:00:00Z",
		moments[0].UTC().Format(time.RFC3339),
	)
}

func TestLabelValuesEndRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(constant.LabelValues, map[string]any{"label": "job"})
	moments := m.Moments()
	assert.Count(t, 2, moments)
	assert.String(
		t,
		"2026-08-16T20:00:00Z",
		moments[1].UTC().Format(time.RFC3339),
	)
}

func TestQueryRangeRunsOnServiceClock(t *testing.T) {
	c, m := clockedInstance(t)
	c.MustCallTool(
		constant.QueryRange,
		map[string]any{
			generativeConstant.ParameterQuery: "up",
			"start":                           "1h",
		},
	)
	moments := m.Moments()
	assert.Count(t, 2, moments)
	assert.String(
		t,
		"2026-08-16T19:00:00Z",
		moments[0].UTC().Format(time.RFC3339),
	)
	assert.String(
		t,
		"2026-08-16T20:00:00Z",
		moments[1].UTC().Format(time.RFC3339),
	)
}
