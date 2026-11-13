package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func reportFixing(
	t *testing.T,
	root string,
	before func(),
) *output.Results {
	t.Helper()
	o := lint_tester.Fixing()
	o.BeforeCommit = before
	var r *output.Results
	testutil.CaptureStdout(func() { r = lint.Report("golint", root, o) })

	return r
}
