package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNormalize(t *testing.T) {
	assert.String(
		t,
		"doc/ai/runbook/lint.md",
		Normalize("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md"),
	)
	assert.String(t, "doc/ai/spec", Normalize("doc/ai/spec/"))
	assert.String(
		t,
		"doc/ai/spec/naming.md",
		Normalize("./doc/ai/spec/naming.md"),
	)
	assert.String(t, "doc/plugins.md", Normalize("doc/plugins.md#convert"))
	assert.String(t, "pkg/lint", Normalize("pkg/lint"))
}
