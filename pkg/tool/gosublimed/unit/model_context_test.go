package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"testing"
)

func TestModelContext(t *testing.T) {
	c := base.New(t).Tools()
	assert.Count(t, 7, c.ListTools())
	created := c.MustCallTool(
		constant.CreateView,
		map[string]any{
			"title":   "test",
			"content": "hello world",
			"syntax":  "Go",
		},
	)
	assert.StringContains(t, "test", created)
	assert.StringContains(t, "hello world", created)
	assert.StringContains(
		t,
		"test",
		c.MustCallTool(constant.ListViews, map[string]any{}),
	)
	assert.StringContains(
		t,
		"hello world",
		c.MustCallTool(constant.ReadView, map[string]any{"view_id": 1}),
	)
	assert.StringContains(
		t,
		"goodbye world",
		c.MustCallTool(
			constant.EditView,
			map[string]any{
				"view_id":    1,
				"old_string": "hello",
				"new_string": "goodbye",
			},
		),
	)
	assert.StringContains(
		t,
		"/tmp/test.go",
		c.MustCallTool(
			constant.OpenFile,
			map[string]any{"file_path": "/tmp/test.go"},
		),
	)
	assert.StringContains(
		t,
		"saved",
		c.MustCallTool(
			constant.SaveView,
			map[string]any{"view_id": 1, "file_path": "/tmp/saved.go"},
		),
	)
	assert.StringContains(
		t,
		"closed",
		c.MustCallTool(constant.CloseView, map[string]any{"view_id": 2}),
	)
}
