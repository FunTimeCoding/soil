package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIsPath(t *testing.T) {
	assert.True(t, IsPath("doc/ai/spec/naming.md"))
	assert.True(t, IsPath(".claude/skills/"))
	assert.True(t, IsPath("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md"))
	assert.True(t, IsPath("../github/soil/pkg"))
	assert.True(t, IsPath("/Users/example/file.md"))
	assert.True(t, IsPath("https://example.org/page"))
	assert.False(t, IsPath(""))
	assert.False(t, IsPath("word"))
	assert.False(t, IsPath("task lint"))
	assert.False(t, IsPath("go run cmd/example/main.go"))
}
