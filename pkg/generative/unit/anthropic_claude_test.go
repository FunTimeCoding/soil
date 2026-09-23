package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/unit/claude_tester"
	"testing"
)

func TestPeekPairsReplyWithItsUserMessage(t *testing.T) {
	c := claude_tester.WriteFixture(
		t,
		[]string{
			claude_tester.UserLine(
				"this is the first user message of the session",
			),
			claude_tester.AssistantLine("reply to the first message"),
			claude_tester.UserLine(
				"this is the second user message of the session",
			),
			claude_tester.AssistantLine(
				"closing reply after the final user message",
			),
		},
	)
	p := c.Peek("fixture")
	assert.Count(t, 2, p.Entries)
	assert.String(
		t,
		"reply to the first message",
		p.Entries[0].AssistantContext,
	)
	assert.String(
		t,
		"closing reply after the final user message",
		p.Entries[1].AssistantContext,
	)
}

func TestPeekWithoutTrailingReply(t *testing.T) {
	c := claude_tester.WriteFixture(
		t,
		[]string{
			claude_tester.UserLine(
				"this is the first user message of the session",
			),
			claude_tester.AssistantLine("reply to the first message"),
			claude_tester.UserLine("this is a final message that got no reply"),
		},
	)
	p := c.Peek("fixture")
	assert.Count(t, 2, p.Entries)
	assert.String(
		t,
		"reply to the first message",
		p.Entries[0].AssistantContext,
	)
	assert.String(t, "", p.Entries[1].AssistantContext)
}

func TestPeekMissingFile(t *testing.T) {
	c := claude.NewDirectory(t.TempDir())
	p := c.Peek("absent")
	assert.Count(t, 0, p.Entries)
}
