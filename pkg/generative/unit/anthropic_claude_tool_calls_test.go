package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/unit/claude_tester"
	"testing"
)

func TestToolCallsCarryThePermissionDecision(t *testing.T) {
	c := claude_tester.WriteFixture(
		t,
		[]string{
			claude_tester.ToolUseLine("toolu_a", "cmp a b", "/repo"),
			claude_tester.RequestBlobLine("toolu_a"),
			claude_tester.ToolResultLine(
				"toolu_a",
				`{"decision":"accept","source":"user_temporary","grantScope":"once","reasonCode":"bashMiss:too-complex"}`,
			),
		},
	)
	calls := c.ToolCalls("fixture")
	assert.Count(t, 1, calls)
	assert.String(t, "cmp a b", calls[0].Detail)
	assert.String(t, "/repo", calls[0].Directory)
	assert.String(t, "accept", calls[0].PermissionDecision)
	assert.String(t, "user_temporary", calls[0].PermissionSource)
	assert.String(t, "bashMiss:too-complex", calls[0].PermissionReason)
}

func TestToolCallsFallBackToTheReasonType(t *testing.T) {
	c := claude_tester.WriteFixture(
		t,
		[]string{
			claude_tester.ToolUseLine("toolu_a", "ls", "/repo"),
			claude_tester.ToolResultLine(
				"toolu_a",
				`{"decision":"accept","source":"config","reasonType":"rule"}`,
			),
		},
	)
	calls := c.ToolCalls("fixture")
	assert.String(t, "config", calls[0].PermissionSource)
	assert.String(t, "rule", calls[0].PermissionReason)
}

func TestToolCallsKeepTheFirstDecisionForACall(t *testing.T) {
	c := claude_tester.WriteFixture(
		t,
		[]string{
			claude_tester.ToolUseLine("toolu_a", "rm x", "/repo"),
			claude_tester.ToolUseLine("toolu_b", "ls", "/repo"),
			claude_tester.ToolUseLine("toolu_c", "pwd", "/repo"),
			claude_tester.ToolResultLine(
				"toolu_b",
				`{"decision":"accept","source":"hook"}`,
			),
			claude_tester.ToolResultLine(
				"toolu_a",
				`{"decision":"reject","source":"user_reject"}`,
			),
			claude_tester.ToolResultLine(
				"toolu_a",
				`{"decision":"accept","source":"config"}`,
			),
			claude_tester.ToolResultLine("toolu_c", "null"),
		},
	)
	calls := c.ToolCalls("fixture")
	assert.Count(t, 3, calls)
	assert.String(t, "user_reject", calls[0].PermissionSource)
	assert.String(t, "reject", calls[0].PermissionDecision)
	assert.String(t, "hook", calls[1].PermissionSource)
	assert.String(t, "", calls[2].PermissionSource)
}
