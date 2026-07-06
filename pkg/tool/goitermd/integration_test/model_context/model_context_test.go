//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/integration_test/model_context_tester"
	"testing"
)

func TestModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	assert.Count(t, 8, o.Client.ListTools())
	n := session.Stub()
	n.Identifier = "sess-1"
	n.TabIdentifier = "tab-1"
	n.TabTitle = "build"
	n.JobName = "zsh"
	o.MockClient.AddSession(n)
	assert.StringContains(
		t,
		"build",
		o.Client.MustCallTool(constant.ListSessions, map[string]any{}),
	)
	assert.StringContains(
		t,
		"session_id",
		o.Client.MustCallTool(
			constant.ReadScreen,
			map[string]any{"session_id": "sess-1"},
		),
	)
	assert.StringContains(
		t,
		"sess-1",
		o.Client.MustCallTool(
			constant.ReadHistory,
			map[string]any{"session_id": "sess-1", "count": 100},
		),
	)
	assert.StringContains(
		t,
		"sent",
		o.Client.MustCallTool(
			constant.SendText,
			map[string]any{"session_id": "sess-1", "text": "ls -la"},
		),
	)
	assert.StringContains(
		t,
		"sent keys",
		o.Client.MustCallTool(
			constant.SendKey,
			map[string]any{
				"session_id": "sess-1",
				"keys":       []string{"Return"},
			},
		),
	)
	assert.StringContains(
		t,
		"set",
		o.Client.MustCallTool(
			constant.SetTabColor,
			map[string]any{
				"session_id": "sess-1",
				"red":        255,
				"green":      0,
				"blue":       0,
			},
		),
	)
	assert.StringContains(
		t,
		"set",
		o.Client.MustCallTool(
			constant.SetTabTitle,
			map[string]any{"tab_id": "tab-1", "title": "deploy"},
		),
	)
	assert.StringContains(
		t,
		"session_id",
		o.Client.MustCallTool(constant.CreateTab, map[string]any{}),
	)
}
