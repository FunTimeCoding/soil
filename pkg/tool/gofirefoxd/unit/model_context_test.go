package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/base"
	"testing"
)

func TestModelContextTabs(t *testing.T) {
	s := base.New(t)
	s.SeedTab()
	c := s.Tools()
	assert.Count(t, 16, c.ListTools())
	assert.StringContains(
		t,
		"Example",
		c.MustCallTool(constant.ListTabs, map[string]any{}),
	)
	assert.StringContains(
		t,
		"example.com",
		c.MustCallTool(constant.ReadTab, map[string]any{"tab_id": 100}),
	)
	assert.StringContains(
		t,
		"tab_id",
		c.MustCallTool(
			constant.CreateTab,
			map[string]any{"url": "https://test.com"},
		),
	)
	assert.StringContains(
		t,
		"closed",
		c.MustCallTool(constant.CloseTab, map[string]any{"tab_id": 100}),
	)
}

func TestModelContextNavigation(t *testing.T) {
	s := base.New(t)
	s.SeedTab()
	c := s.Tools()
	assert.StringContains(
		t,
		"navigated to",
		c.MustCallTool(
			constant.Navigate,
			map[string]any{"tab_id": 100, "url": "https://other.com"},
		),
	)
	assert.StringContains(
		t,
		"navigated back",
		c.MustCallTool(constant.NavigateBack, map[string]any{"tab_id": 100}),
	)
	assert.StringContains(
		t,
		"tab_id",
		c.MustCallTool(
			constant.WaitFor,
			map[string]any{"tab_id": 100, "selector": "#results"},
		),
	)
}

func TestModelContextGroups(t *testing.T) {
	s := base.New(t)
	s.SeedTab()
	c := s.Tools()
	assert.StringContains(
		t,
		"grouped into group",
		c.MustCallTool(
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
		c.MustCallTool(
			constant.UpdateGroup,
			map[string]any{"group_id": 1, "title": "deploy"},
		),
	)
	assert.StringContains(
		t,
		"ungrouped",
		c.MustCallTool(constant.UngroupTab, map[string]any{"tab_id": 100}),
	)
}

func TestModelContextPage(t *testing.T) {
	s := base.New(t)
	s.SeedTab()
	c := s.Tools()
	assert.StringContains(
		t,
		"document.title",
		c.MustCallTool(
			constant.Evaluate,
			map[string]any{"tab_id": 100, "expression": "document.title"},
		),
	)
	assert.StringContains(
		t,
		"#submit",
		c.MustCallTool(
			constant.Click,
			map[string]any{"tab_id": 100, "selector": "#submit"},
		),
	)
	assert.StringContains(
		t,
		"operator",
		c.MustCallTool(
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
		c.MustCallTool(constant.Screenshot, map[string]any{"tab_id": 100}),
	)
}

func TestModelContextDownloads(t *testing.T) {
	s := base.New(t)
	s.SeedTab()
	c := s.Tools()
	assert.StringContains(
		t,
		"download_id",
		c.MustCallTool(
			constant.DownloadFile,
			map[string]any{"url": "https://example.com/paper.pdf"},
		),
	)
	assert.StringContains(
		t,
		"in_progress",
		c.MustCallTool(constant.DownloadStatus, map[string]any{}),
	)
}
