package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func FixingReport(
	t *testing.T,
	root string,
	before func(),
) *output.Results {
	t.Helper()
	o := Fixing()
	o.BeforeCommit = before
	var r *output.Results
	testutil.CaptureStdout(func() { r = lint.Report("golint", root, o) })

	return r
}
