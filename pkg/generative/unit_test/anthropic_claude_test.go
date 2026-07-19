package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"path/filepath"
	"testing"
)

func writeFixture(t *testing.T, lines []string) *claude.Client {
	t.Helper()
	base := t.TempDir()
	f := system.Create(filepath.Join(base, "fixture.jsonl"))

	for _, line := range lines {
		writer.Print(f, "%s\n", line)
	}

	errors.PanicClose(f)

	return claude.NewDirectory(base)
}

func userLine(text string) string {
	return fmt.Sprintf(
		`{"type":"user","message":{"content":"%s"}}`,
		text,
	)
}

func assistantLine(text string) string {
	return fmt.Sprintf(
		`{"type":"assistant","message":{"content":[{"type":"text","text":"%s"}]}}`,
		text,
	)
}

func TestPeekPairsReplyWithItsUserMessage(t *testing.T) {
	c := writeFixture(
		t,
		[]string{
			userLine("this is the first user message of the session"),
			assistantLine("reply to the first message"),
			userLine("this is the second user message of the session"),
			assistantLine("closing reply after the final user message"),
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
	c := writeFixture(
		t,
		[]string{
			userLine("this is the first user message of the session"),
			assistantLine("reply to the first message"),
			userLine("this is a final message that got no reply"),
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
