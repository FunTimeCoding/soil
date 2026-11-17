package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/base"
	"os"
	"path/filepath"
	"testing"
)

func TestMissingArgumentsAreRefusedBeforeBrowser(t *testing.T) {
	s := base.New(t)
	c := s.Tools()

	for name, arguments := range map[string]map[string]any{
		constant.Navigate:     {"tab_id": 1, "url": ""},
		constant.GroupTabs:    {"tab_ids": []int{}},
		constant.Evaluate:     {"tab_id": 1, "expression": ""},
		constant.Click:        {"tab_id": 1, "selector": ""},
		constant.Fill:         {"tab_id": 1, "selector": ""},
		constant.DownloadFile: {"url": ""},
	} {
		assert.StringContains(
			t,
			"required",
			c.MustCallToolError(name, arguments),
		)
	}

	assert.Count(t, 0, s.Reporter.Events())
	tabs, e := s.Source.Tabs()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, tabs)
}

func TestScreenshotIsSavedToDownloadDirectory(t *testing.T) {
	s := base.New(t)
	b := tab.New()
	b.Identifier = 3
	b.Title = "Example"
	s.Source.AddTab(b)
	assert.StringContains(
		t,
		"firefox_3.png",
		s.Tools().MustCallTool(constant.Screenshot, map[string]any{"tab_id": 3}),
	)
	content, e := os.ReadFile(filepath.Join(s.Directory, "firefox_3.png"))
	assert.FatalOnError(t, e)
	assert.String(t, "mock image", string(content))
}
