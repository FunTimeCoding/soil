package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_change"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCoverageComputeWindows(t *testing.T) {
	servers := coverage.Compute(
		[]*tool_call.Call{
			coverageCall("mcp__alfa__list_items", "2026-08-20T10:00:00Z"),
			coverageCall("mcp__alfa__list_items", "2026-01-05T10:00:00Z"),
			coverageCall("mcp__alfa__get_item", "2026-01-05T10:00:00Z"),
		},
		map[string][]string{"alfa": {"list_items", "get_item", "delete_item"}},
		map[string]string{"alfa": "pkg/tool/goalfad"},
		map[string]bool{"alfa": true},
		nil,
		coverageNow(),
	)
	assert.Integer(t, 1, len(servers))
	s := servers[0]
	assert.String(t, "alfa", s.Name)
	assert.String(t, "pkg/tool/goalfad", s.Path)
	assert.True(t, s.Configured)
	assert.Integer(t, 3, s.Registered)
	assert.Integer(t, 2, s.UsedTotal)
	assert.Integer(t, 1, s.UsedRecent)
	assert.Integer(t, 3, s.CallsTotal)
	assert.Integer(t, 1, s.CallsRecent)
	assert.Integer(t, 3, len(s.Tools))
	assert.String(t, "list_items", s.Tools[0].Name)
	assert.Integer(t, 2, s.Tools[0].CallsTotal)
	assert.Integer(t, 1, s.Tools[0].CallsRecent)
}

func TestCoverageComputeRetiredName(t *testing.T) {
	servers := coverage.Compute(
		[]*tool_call.Call{
			coverageCall("mcp__alfa__old_name", "2026-08-20T10:00:00Z"),
		},
		map[string][]string{"alfa": {"new_name"}},
		map[string]string{"alfa": "pkg/tool/goalfad"},
		map[string]bool{"alfa": true},
		nil,
		coverageNow(),
	)
	s := servers[0]
	assert.Integer(t, 1, s.Registered)
	assert.Integer(t, 0, s.UsedTotal)
	assert.Integer(t, 1, s.CallsTotal)
	assert.Integer(t, 2, len(s.Tools))
	assert.String(t, "old_name", s.Tools[0].Name)
	assert.False(t, s.Tools[0].Registered)
}

func TestCoverageComputeConfiguredOnly(t *testing.T) {
	servers := coverage.Compute(
		nil,
		map[string][]string{},
		map[string]string{},
		map[string]bool{"github": true},
		nil,
		coverageNow(),
	)
	assert.Integer(t, 1, len(servers))
	assert.String(t, "github", servers[0].Name)
	assert.Integer(t, 0, servers[0].Registered)
	assert.True(t, servers[0].Configured)
}

func TestCoverageComputeIgnoresOtherTools(t *testing.T) {
	servers := coverage.Compute(
		[]*tool_call.Call{
			coverageCall("Bash", "2026-08-20T10:00:00Z"),
			coverageCall("Read", "2026-08-20T10:00:00Z"),
		},
		map[string][]string{},
		map[string]string{},
		map[string]bool{},
		nil,
		coverageNow(),
	)
	assert.Integer(t, 0, len(servers))
}

func TestCoverageComputeAliasFold(t *testing.T) {
	servers := coverage.Compute(
		[]*tool_call.Call{
			coverageCall("mcp__alfa__list_items", "2026-08-20T10:00:00Z"),
			coverageCall(
				"mcp__claude_ai_Alfa__list_items",
				"2026-08-20T11:00:00Z",
			),
			coverageCall("mcp__bravo__inspect", "2026-08-20T10:00:00Z"),
		},
		map[string][]string{"alfa": {"list_items"}},
		map[string]string{"alfa": "pkg/tool/goalfad"},
		map[string]bool{"alfa": true},
		map[string]string{"claude_ai_Alfa": "alfa"},
		coverageNow(),
	)
	assert.Integer(t, 2, len(servers))
	assert.String(t, "alfa", servers[0].Name)
	assert.Integer(t, 2, servers[0].CallsTotal)
	assert.Integer(t, 1, len(servers[0].Tools))
	assert.Integer(t, 2, servers[0].Tools[0].CallsTotal)
	assert.String(t, "bravo", servers[1].Name)
}

func TestCoverageComputeUnconfiguredObserved(t *testing.T) {
	servers := coverage.Compute(
		[]*tool_call.Call{
			coverageCall("mcp__gone__old_tool", "2026-08-20T10:00:00Z"),
		},
		map[string][]string{},
		map[string]string{},
		map[string]bool{},
		nil,
		coverageNow(),
	)
	assert.Integer(t, 1, len(servers))
	assert.String(t, "gone", servers[0].Name)
	assert.False(t, servers[0].Configured)
	assert.Integer(t, 0, servers[0].Registered)
	assert.Integer(t, 1, servers[0].CallsTotal)
}

func TestLabelChangeFormatFromUnset(t *testing.T) {
	assert.String(
		t,
		"colour (unset)→green",
		label_change.Format("colour", "", "green"),
	)
}

func TestLabelChangeFormatReplace(t *testing.T) {
	assert.String(
		t,
		"colour blue→green",
		label_change.Format("colour", "blue", "green"),
	)
}

func TestLabelChangeFormatRemove(t *testing.T) {
	assert.String(
		t,
		"colour blue→ (unset)",
		label_change.Format("colour", "blue", ""),
	)
}

func TestTranscriptCacheParity(t *testing.T) {
	base := t.TempDir()
	system.WriteFile(
		filepath.Join(base, "alfa.jsonl"),
		[]byte(transcriptLine("Bash", "2026-08-20T10:00:00Z")),
		0644,
	)
	c := transcript_cache.New(claude.NewDirectory(base))
	sessions := c.Sessions()
	assert.Integer(t, 1, len(sessions))
	assert.String(t, "alfa", sessions[0].Identifier)
	assert.Integer(t, 1, len(c.ToolCalls("alfa")))
	assert.Integer(t, 1, len(c.ToolCalls("alfa")))
	assert.Integer(t, 1, len(c.Sessions()))
}

func TestTranscriptCacheGrownFile(t *testing.T) {
	base := t.TempDir()
	path := filepath.Join(base, "alfa.jsonl")
	system.WriteFile(
		path,
		[]byte(transcriptLine("Bash", "2026-08-20T10:00:00Z")),
		0644,
	)
	c := transcript_cache.New(claude.NewDirectory(base))
	assert.Integer(t, 1, len(c.ToolCalls("alfa")))
	f, e := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	assert.Nil(t, e)
	_, e = f.WriteString(transcriptLine("Read", "2026-08-20T11:00:00Z"))
	assert.Nil(t, e)
	errors.PanicClose(f)
	assert.Integer(t, 2, len(c.ToolCalls("alfa")))
}

func TestTranscriptCacheDelete(t *testing.T) {
	base := t.TempDir()
	system.WriteFile(
		filepath.Join(base, "alfa.jsonl"),
		[]byte(transcriptLine("Bash", "2026-08-20T10:00:00Z")),
		0644,
	)
	c := transcript_cache.New(claude.NewDirectory(base))
	assert.Integer(t, 1, len(c.ToolCalls("alfa")))
	c.Delete("alfa")
	assert.Integer(t, 0, len(c.ToolCalls("alfa")))
	assert.Integer(t, 0, len(c.Sessions()))
}

func TestFiveHourResetTextMinutes(t *testing.T) {
	assert.String(
		t,
		"in 33 min",
		result(time.Now().Add(33*time.Minute+time.Second)).FiveHourResetText(),
	)
}

func TestFiveHourResetTextHours(t *testing.T) {
	assert.String(
		t,
		"in 2 hr 5 min",
		result(
			time.Now().Add(2*time.Hour+5*time.Minute+time.Second),
		).FiveHourResetText(),
	)
}

func TestFiveHourResetTextPassed(t *testing.T) {
	assert.String(
		t,
		"now",
		result(time.Now().Add(-time.Minute)).FiveHourResetText(),
	)
}

func TestSevenDayResetText(t *testing.T) {
	assert.String(t, "Wed 21:00", result(time.Now()).SevenDayResetText())
}

func TestHasFable(t *testing.T) {
	assert.True(t, result(time.Now()).HasFable())
}

func TestFableResetTextFallsBackToText(t *testing.T) {
	assert.String(t, "Wed 8:59 PM", result(time.Now()).FableResetText())
}

func TestFableResetTextPrefersTimestamp(t *testing.T) {
	r := usage_result.New(
		26,
		time.Now(),
		20,
		time.Now(),
		34,
		"",
		time.Date(2026, 9, 15, 16, 59, 59, 0, time.Local),
		time.Now(),
	)
	assert.True(t, r.HasFable())
	assert.String(t, "Tue 16:59", r.FableResetText())
}
