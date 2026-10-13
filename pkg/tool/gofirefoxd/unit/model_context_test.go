package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	assert.Count(t, 16, o.Client.ListTools())
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
		o.Client.MustCallTool(constant.ReadTab, map[string]any{"tab_id": 100}),
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
		"tab_id",
		o.Client.MustCallTool(
			constant.WaitFor,
			map[string]any{"tab_id": 100, "selector": "#results"},
		),
	)
	assert.StringContains(
		t,
		"document.title",
		o.Client.MustCallTool(
			constant.Evaluate,
			map[string]any{"tab_id": 100, "expression": "document.title"},
		),
	)
	assert.StringContains(
		t,
		"#submit",
		o.Client.MustCallTool(
			constant.Click,
			map[string]any{"tab_id": 100, "selector": "#submit"},
		),
	)
	assert.StringContains(
		t,
		"operator",
		o.Client.MustCallTool(
			constant.Fill,
			map[string]any{
				"tab_id":   100,
				"selector": "#user",
				"value":    "operator",
			},
		),
	)
	assert.StringContains(
		t,
		"firefox_100.png",
		o.Client.MustCallTool(
			constant.Screenshot,
			map[string]any{"tab_id": 100},
		),
	)
	assert.StringContains(
		t,
		"download_id",
		o.Client.MustCallTool(
			constant.DownloadFile,
			map[string]any{"url": "https://example.com/paper.pdf"},
		),
	)
	assert.StringContains(
		t,
		"in_progress",
		o.Client.MustCallTool(constant.DownloadStatus, map[string]any{}),
	)
	assert.StringContains(
		t,
		"closed",
		o.Client.MustCallTool(constant.CloseTab, map[string]any{"tab_id": 100}),
	)
}
