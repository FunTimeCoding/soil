package integration_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	c := o.Client
	assert.Count(t, 4, c.ListTools())
	c.MustCallTool(
		constant.AddEntry,
		map[string]any{
			constant.Action:      "restarted web server",
			constant.User:        "jdoe",
			constant.System:      "worker1",
			constant.Service:     "nginx",
			constant.Description: "nginx was unresponsive",
		},
	)
	assert.StringContains(
		t,
		"1 entries",
		c.MustCallTool(constant.ListEntries, map[string]any{}),
	)
	assert.StringContains(
		t,
		"1 entries",
		c.MustCallTool(
			constant.ListEntries,
			map[string]any{constant.System: "worker1"},
		),
	)
	assert.StringContains(
		t,
		"1 entries",
		c.MustCallTool(
			constant.ListEntries,
			map[string]any{constant.User: "jdoe"},
		),
	)
	assert.StringContains(
		t,
		"0 entries",
		c.MustCallTool(
			constant.ListEntries,
			map[string]any{constant.System: "nonexistent"},
		),
	)
	assert.StringContains(
		t,
		"updated successfully",
		c.MustCallTool(
			constant.UpdateEntry,
			map[string]any{
				constant.Identifier: float64(1),
				constant.Action:     "cleared and documented",
			},
		),
	)
	assert.StringContains(
		t,
		"cleared and documented",
		c.MustCallTool(constant.ListEntries, map[string]any{}),
	)
	assert.StringContains(
		t,
		"deleted successfully",
		c.MustCallTool(
			constant.DeleteEntry,
			map[string]any{constant.Identifier: float64(1)},
		),
	)
	assert.StringContains(
		t,
		"0 entries",
		c.MustCallTool(constant.ListEntries, map[string]any{}),
	)
}
