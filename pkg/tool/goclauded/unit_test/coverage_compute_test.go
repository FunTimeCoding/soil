package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"testing"
	"time"
)

func coverageNow() time.Time {
	return time.Date(2026, 8, 22, 12, 0, 0, 0, time.UTC)
}

func coverageCall(
	name string,
	timestamp string,
) tool_call.Call {
	return tool_call.Call{Name: name, Timestamp: timestamp}
}

func TestCoverageComputeWindows(t *testing.T) {
	servers := coverage.Compute(
		[]tool_call.Call{
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
		[]tool_call.Call{
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
		[]tool_call.Call{
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
		[]tool_call.Call{
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
		[]tool_call.Call{
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
