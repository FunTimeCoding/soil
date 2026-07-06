//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	c := o.Client
	assert.Count(t, 4, c.ListTools())
	status := c.MustCallTool(constant.GetStatus, map[string]any{})
	assert.StringContains(t, "2", status)
	alerts := c.MustCallTool(
		constant.GetAlerts,
		map[string]any{parameter.Name: "HighMemory"},
	)
	assert.StringContains(t, "fp1", alerts)
	assert.StringContains(t, "critical", alerts)
	recent := c.MustCallTool(
		constant.GetRecentAlerts,
		map[string]any{},
	)
	assert.StringContains(t, "HighMemory", recent)
	assert.StringContains(t, "DiskFull", recent)
	top := c.MustCallTool(constant.GetTopAlerts, map[string]any{})
	assert.StringContains(t, "HighMemory", top)
}
