package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache"
	"os"
	"path/filepath"
	"testing"
)

func transcriptLine(name string, timestamp string) string {
	return fmt.Sprintf(
		`{"type":"assistant","timestamp":"%s","message":{"content":[{"type":"tool_use","id":"t1","name":"%s","input":{}}]}}%s`,
		timestamp,
		name,
		"\n",
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
