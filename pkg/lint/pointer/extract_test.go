package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	assert.Any(
		t,
		[]string{"doc/ai/spec/naming.md"},
		Extract("Read `doc/ai/spec/naming.md` before."),
	)
	assert.Any(
		t,
		[]string{"pkg/lint", "doc/ai"},
		Extract("See `pkg/lint` and `doc/ai`."),
	)
	assert.Any(
		t,
		[]string{"doc/ai/runbook/lint.md"},
		Extract("[the runbook](doc/ai/runbook/lint.md)"),
	)
	assert.Any(t, []string(nil), Extract("Run `task lint` now."))
	assert.Any(t, []string(nil), Extract("stray ` backtick"))
	assert.Any(t, []string(nil), Extract(""))
}
