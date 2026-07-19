package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestRelative(t *testing.T) {
	result, inside := pointer.Relative(
		"doc/ai/runbook/lint.md",
		"../spec/naming.md",
	)
	assert.String(t, "doc/ai/spec/naming.md", result)
	assert.True(t, inside)
	result, inside = pointer.Relative(
		"doc/ai/spec/error-handling/README.md",
		"../../runbook/lint.md",
	)
	assert.String(t, "doc/ai/runbook/lint.md", result)
	assert.True(t, inside)
	result, inside = pointer.Relative("CLAUDE.md", "../sibling/doc/README.md")
	assert.String(t, "../sibling/doc/README.md", result)
	assert.False(t, inside)
}
