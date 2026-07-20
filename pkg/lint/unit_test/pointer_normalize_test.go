package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestNormalize(t *testing.T) {
	assert.String(
		t,
		"doc/ai/runbook/lint.md",
		pointer.Normalize("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md"),
	)
	assert.String(t, "doc/ai/spec", pointer.Normalize("doc/ai/spec/"))
	assert.String(
		t,
		"doc/ai/spec/naming.md",
		pointer.Normalize("./doc/ai/spec/naming.md"),
	)
	assert.String(
		t,
		"doc/plugins.md",
		pointer.Normalize("doc/plugins.md#convert"),
	)
	assert.String(t, "pkg/lint", pointer.Normalize("pkg/lint"))
}
