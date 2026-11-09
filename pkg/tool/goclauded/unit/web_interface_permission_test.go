package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"strings"
	"testing"
	"time"
)

func TestPermissionsPageRanksBlockerWithItsCommands(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	e := session.New("session-a")
	e.WorkDirectory = project
	s.Set.Client.AddSession(e)
	c := tool_call.New(
		"Bash",
		"toolu_a",
		time.Now().Add(-time.Hour).Format(time.RFC3339Nano),
	)
	c.Detail = "sort x"
	c.PermissionDecision = "accept"
	c.PermissionSource = "user_temporary"
	c.PermissionReason = "bashMiss:no-rule-match"
	s.Set.Client.AddToolCall("session-a", c)
	page := body(t, s.Server, constant.PermissionsPath)
	markup_tester.AssertCount(t, 1, page, "tbody tr")
	markup_tester.AssertText(
		t,
		"no-rule-match",
		page,
		"tbody tr td:nth-child(4)",
	)
	markup_tester.AssertText(t, "sort", page, "details summary")
	markup_tester.AssertText(t, "sort x", page, "details pre")
	assert.StringContains(t, "1 prompts", page)
}

func TestPermissionsPageCutsLongCommand(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	project := t.TempDir()
	permissionSettings(t, project, "settings.local.json", "Bash(ls:*)")
	e := session.New("session-a")
	e.WorkDirectory = project
	s.Set.Client.AddSession(e)
	c := tool_call.New(
		"Bash",
		"toolu_a",
		time.Now().Add(-time.Hour).Format(time.RFC3339Nano),
	)
	c.Detail = join.Empty("sort ", strings.Repeat("a", 500))
	c.PermissionSource = "user_temporary"
	s.Set.Client.AddToolCall("session-a", c)
	page := body(t, s.Server, constant.PermissionsPath)
	markup_tester.AssertText(
		t,
		join.Empty("sort ", strings.Repeat("a", 395), "…"),
		page,
		"details pre",
	)
}

func TestPermissionsPageSaysWhenNothingPrompted(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	page := body(t, s.Server, constant.PermissionsPath)
	assert.StringContains(t, "No permission prompts in this window.", page)
	markup_tester.AssertCount(t, 0, page, "details")
}
