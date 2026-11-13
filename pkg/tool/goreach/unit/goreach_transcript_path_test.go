package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goreach"
	"testing"
)

func TestTranscriptPathWithoutSessionIsNotSelected(t *testing.T) {
	t.Setenv(goclaudeConstant.SessionIdentifierEnvironment, "")
	_, e := goreach.TranscriptPath("", nil)
	assert.True(t, not_selected.Is(e))
}

func TestTranscriptPathWithoutTranscriptIsNotFound(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	t.Setenv(goclaudeConstant.SessionIdentifierEnvironment, "alfa-bravo")
	_, e := goreach.TranscriptPath("", nil)
	assert.True(t, not_found.Is(e))
}

func TestTranscriptPathTakesFlagAsGiven(t *testing.T) {
	path, e := goreach.TranscriptPath("charlie.jsonl", nil)
	assert.FatalOnError(t, e)
	assert.String(t, "charlie.jsonl", path)
}
