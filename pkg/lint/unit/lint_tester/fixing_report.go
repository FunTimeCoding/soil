package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func (o *Tester) FixingReport(
	root string,
	before func(),
) *output.Results {
	o.t.Helper()
	p := Fixing()
	p.BeforeCommit = before
	var r *output.Results
	testutil.CaptureStdout(func() { r = lint.Report("golint", root, p) })

	return r
}
