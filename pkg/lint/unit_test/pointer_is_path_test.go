package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestIsPath(t *testing.T) {
	assert.True(t, pointer.IsPath("doc/ai/spec/naming.md"))
	assert.True(t, pointer.IsPath(".claude/skills/"))
	assert.True(t, pointer.IsPath("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md"))
	assert.True(t, pointer.IsPath("../github/soil/pkg"))
	assert.True(t, pointer.IsPath("/Users/example/file.md"))
	assert.True(t, pointer.IsPath("https://example.org/page"))
	assert.False(t, pointer.IsPath(""))
	assert.False(t, pointer.IsPath("word"))
	assert.False(t, pointer.IsPath("task lint"))
	assert.False(t, pointer.IsPath("go run cmd/example/main.go"))
}
