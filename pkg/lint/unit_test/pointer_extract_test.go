package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestExtract(t *testing.T) {
	assert.Any(
		t,
		[]string{"doc/ai/spec/naming.md"},
		pointer.Extract("Read `doc/ai/spec/naming.md` before."),
	)
	assert.Any(
		t,
		[]string{"pkg/lint", "doc/ai"},
		pointer.Extract("See `pkg/lint` and `doc/ai`."),
	)
	assert.Any(
		t,
		[]string{"doc/ai/runbook/lint.md"},
		pointer.Extract("[the runbook](doc/ai/runbook/lint.md)"),
	)
	assert.Any(t, []string(nil), pointer.Extract("Run `task lint` now."))
	assert.Any(t, []string(nil), pointer.Extract("stray ` backtick"))
	assert.Any(t, []string(nil), pointer.Extract(""))
}
