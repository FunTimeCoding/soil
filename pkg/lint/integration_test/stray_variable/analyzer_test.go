package stray_variable

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/stray_variable"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagged(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	stray_variable.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	stray_variable.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
