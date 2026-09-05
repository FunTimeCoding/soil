package expected_first

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/expected_first"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagged(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	expected_first.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
}

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	expected_first.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
