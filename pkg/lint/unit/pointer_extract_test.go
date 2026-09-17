package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestExtract(t *testing.T) {
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewSpan("doc/ai/spec/naming.md")},
		pointer.Extract("Read `doc/ai/spec/naming.md` before."),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{
			pointer.NewSpan("pkg/lint"),
			pointer.NewSpan("doc/ai"),
		},
		pointer.Extract("See `pkg/lint` and `doc/ai`."),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewLink("doc/ai/runbook/lint.md")},
		pointer.Extract("[the runbook](doc/ai/runbook/lint.md)"),
	)
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewLink("naming.md")},
		pointer.Extract("[naming](naming.md)"),
	)
	assert.Any(t, []*pointer.Candidate(nil), pointer.Extract("[top](#top)"))
	assert.Any(
		t,
		[]*pointer.Candidate{pointer.NewLink("https://host.example")},
		pointer.Extract("[site](https://host.example)"),
	)
	assert.Any(t, []*pointer.Candidate(nil), pointer.Extract("[bare](naming)"))
	assert.Any(
		t,
		[]*pointer.Candidate(nil),
		pointer.Extract("Run `task lint` now."),
	)
	assert.Any(
		t,
		[]*pointer.Candidate(nil),
		pointer.Extract("stray ` backtick"),
	)
	assert.Any(t, []*pointer.Candidate(nil), pointer.Extract(""))
}
