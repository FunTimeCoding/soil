package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 7, o.Client.ListTools())
	created := o.Client.MustCallTool(
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
		o.Client.MustCallTool(constant.ListViews, map[string]any{}),
	)
	assert.StringContains(
		t,
		"hello world",
		o.Client.MustCallTool(
			constant.ReadView,
			map[string]any{"view_id": 1},
		),
	)
	assert.StringContains(
		t,
		"goodbye world",
		o.Client.MustCallTool(
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
		o.Client.MustCallTool(
			constant.OpenFile,
			map[string]any{"file_path": "/tmp/test.go"},
		),
	)
	assert.StringContains(
		t,
		"saved",
		o.Client.MustCallTool(
			constant.SaveView,
			map[string]any{"view_id": 1, "file_path": "/tmp/saved.go"},
		),
	)
	assert.StringContains(
		t,
		"closed",
		o.Client.MustCallTool(
			constant.CloseView,
			map[string]any{"view_id": 2},
		),
	)
}
