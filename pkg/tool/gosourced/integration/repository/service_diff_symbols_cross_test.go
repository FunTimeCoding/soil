package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/report_tester"
	"testing"
)

func TestDiffSymbolsFollowsMoveIntoRequiredModule(t *testing.T) {
	v := crossService(t)
	r, e := v.Diff.DiffSymbols("user", "HEAD~1", "HEAD")
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{
			"unchanged example.org/user/echo echo Echo",
			"moved example.org/lib/charlie charlie Bravo",
		},
		report_tester.Lines(r),
	)
	assert.Integer(t, 1, len(r.Modules))
	assert.String(t, "lib", r.Modules[0].Name)
	assert.String(t, "v0.1.0", r.Modules[0].From)
	assert.String(t, "v0.2.0", r.Modules[0].To)
}
