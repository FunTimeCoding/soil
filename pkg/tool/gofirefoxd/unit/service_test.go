package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/model_context_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestTheServiceClosesItsClient(t *testing.T) {
	s := base.New(t)
	s.Service.Close()
	assert.Integer(t, 1, s.MockClient.Closed)
}

func TestMissingArgumentsAreRefusedBeforeTheBrowser(t *testing.T) {
	o := model_context_tester.New(t)

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
			o.Client.MustCallToolError(name, arguments),
		)
	}

	assert.Count(t, 0, o.Reporter.Events())
	tabs, e := o.MockClient.Tabs()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, tabs)
}

func TestAScreenshotIsSavedToTheDownloadDirectory(t *testing.T) {
	o := model_context_tester.New(t)
	b := tab.New()
	b.Identifier = 3
	b.Title = "Example"
	o.MockClient.AddTab(b)
	assert.StringContains(
		t,
		"firefox_3.png",
		o.Client.MustCallTool(constant.Screenshot, map[string]any{"tab_id": 3}),
	)
	content, e := os.ReadFile(filepath.Join(o.Directory, "firefox_3.png"))
	assert.FatalOnError(t, e)
	assert.String(t, "mock image", string(content))
}
