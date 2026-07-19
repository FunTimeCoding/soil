package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 9, o.Client.ListTools())
	b := tab.New()
	b.Identifier = 100
	b.Locator = "https://example.com"
	b.Title = "Example"
	o.MockClient.AddTab(b)
	assert.StringContains(
		t,
		"Example",
		o.Client.MustCallTool(constant.ListTabs, map[string]any{}),
	)
	assert.StringContains(
		t,
		"example.com",
		o.Client.MustCallTool(
			constant.ReadTab,
			map[string]any{"tab_id": 100},
		),
	)
	assert.StringContains(
		t,
		"tab_id",
		o.Client.MustCallTool(
			constant.CreateTab,
			map[string]any{"url": "https://test.com"},
		),
	)
	assert.StringContains(
		t,
		"navigated to",
		o.Client.MustCallTool(
			constant.Navigate,
			map[string]any{"tab_id": 100, "url": "https://other.com"},
		),
	)
	assert.StringContains(
		t,
		"navigated back",
		o.Client.MustCallTool(
			constant.NavigateBack,
			map[string]any{"tab_id": 100},
		),
	)
	assert.StringContains(
		t,
		"grouped into group",
		o.Client.MustCallTool(
			constant.GroupTabs,
			map[string]any{
				"tab_ids": []int{100},
				"title":   "work",
				"color":   "blue",
			},
		),
	)
	assert.StringContains(
		t,
		"group updated",
		o.Client.MustCallTool(
			constant.UpdateGroup,
			map[string]any{"group_id": 1, "title": "deploy"},
		),
	)
	assert.StringContains(
		t,
		"ungrouped",
		o.Client.MustCallTool(
			constant.UngroupTab,
			map[string]any{"tab_id": 100},
		),
	)
	assert.StringContains(
		t,
		"closed",
		o.Client.MustCallTool(
			constant.CloseTab,
			map[string]any{"tab_id": 100},
		),
	)
}
