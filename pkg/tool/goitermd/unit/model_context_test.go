package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/iterm/session"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"testing"
)

func TestModelContext(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.Count(t, 8, c.ListTools())
	n := session.Stub()
	n.Identifier = "sess-1"
	n.TabIdentifier = "tab-1"
	n.TabTitle = "build"
	n.JobName = "zsh"
	s.Source.AddSession(n)
	assert.StringContains(
		t,
		"build",
		c.MustCallTool(constant.ListSessions, map[string]any{}),
	)
	assert.StringContains(
		t,
		"session_id",
		c.MustCallTool(
			constant.ReadScreen,
			map[string]any{"session_id": "sess-1"},
		),
	)
	assert.StringContains(
		t,
		"sess-1",
		c.MustCallTool(
			constant.ReadHistory,
			map[string]any{"session_id": "sess-1", "count": 100},
		),
	)
	assert.StringContains(
		t,
		"sent",
		c.MustCallTool(
			constant.SendText,
			map[string]any{"session_id": "sess-1", "text": "ls -la"},
		),
	)
	assert.StringContains(
		t,
		"sent keys",
		c.MustCallTool(
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
		c.MustCallTool(
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
		c.MustCallTool(
			constant.SetTabTitle,
			map[string]any{"tab_id": "tab-1", "title": "deploy"},
		),
	)
	assert.StringContains(
		t,
		"session_id",
		c.MustCallTool(constant.CreateTab, map[string]any{}),
	)
}
