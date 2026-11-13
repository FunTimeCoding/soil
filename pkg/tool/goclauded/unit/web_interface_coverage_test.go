package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"testing"
	"time"
)

func TestCoverageRowOpensItsToolsInPlace(t *testing.T) {
	t.Setenv(constant.CoverageRootEnvironment, t.TempDir())
	s := base.New(t)
	s.Set.Client.AddSession(session.New("session-a"))
	s.Set.Client.AddToolCall(
		"session-a",
		coverageCall(
			"mcp__alfa__status",
			time.Now().Add(-time.Hour).Format(time.RFC3339Nano),
		),
	)
	page := body(t, s.Server, constant.CoveragePath)
	markup_tester.AssertCount(t, 1, page, "tbody tr.opening-row")
	markup_tester.AssertCount(t, 0, page, "table table")
	markup_tester.AssertAttribute(
		t,
		"/coverage/open?server=alfa",
		page,
		"tr.opening-row",
		"hx-get",
	)
	opened := body(t, s.Server, "/coverage/open?server=alfa")
	assert.StringContains(t, `class="opening-detail"><td colspan="6">`, opened)
	markup_tester.AssertText(t, "status", opened, "table td:first-child")
	markup_tester.AssertAttribute(
		t,
		"/coverage/close?server=alfa",
		opened,
		"template tr.opening-row",
		"hx-get",
	)
	closed := body(t, s.Server, "/coverage/close?server=alfa")
	assert.StringContains(t, `hidden="hidden"></tr>`, closed)
	markup_tester.AssertAttribute(
		t,
		"/coverage/open?server=alfa",
		closed,
		"template tr.opening-row",
		"hx-get",
	)
	gone := body(t, s.Server, "/coverage/open?server=gone")
	markup_tester.AssertAttribute(
		t,
		"delete",
		gone,
		"template tr",
		"hx-swap-oob",
	)
}
