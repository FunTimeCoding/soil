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
	markup_tester.AssertCount(t, 1, page, "tbody tr.opening-row")
	markup_tester.AssertText(
		t,
		"no-rule-match",
		page,
		"tbody tr.opening-row td:nth-child(4)",
	)
	markup_tester.AssertCount(t, 0, page, "pre")
	assert.StringContains(t, "1 prompts", page)
	markup_tester.AssertAttribute(
		t,
		"/permissions/open?blocker=sort",
		page,
		"tr.opening-row",
		"hx-get",
	)
	opened := body(t, s.Server, "/permissions/open?blocker=sort")
	assert.StringContains(t, `class="opening-detail"><td colspan="5">`, opened)
	markup_tester.AssertText(t, "sort x", opened, "pre")
	markup_tester.AssertAttribute(
		t,
		"/permissions/close?blocker=sort",
		opened,
		"template tr.opening-row",
		"hx-get",
	)
	closed := body(t, s.Server, "/permissions/close?blocker=sort")
	markup_tester.AssertCount(t, 0, closed, "pre")
	markup_tester.AssertAttribute(
		t,
		"/permissions/open?blocker=sort",
		closed,
		"template tr.opening-row",
		"hx-get",
	)
}

func TestPermissionsRowGoneFromReportLeavesPage(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	gone := body(t, s.Server, "/permissions/open?blocker=sort")
	markup_tester.AssertAttribute(
		t,
		"delete",
		gone,
		"template tr",
		"hx-swap-oob",
	)
	assert.StringContains(t, `<tr id="blocker-`, gone)
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
	markup_tester.AssertText(
		t,
		join.Empty("sort ", strings.Repeat("a", 395), "…"),
		body(t, s.Server, "/permissions/open?blocker=sort"),
		"pre",
	)
}

func TestPermissionsPageSaysWhenNothingPrompted(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	s := base.New(t)
	page := body(t, s.Server, constant.PermissionsPath)
	assert.StringContains(t, "No permission prompts in this window.", page)
	markup_tester.AssertCount(t, 0, page, "tr.opening-row")
}
