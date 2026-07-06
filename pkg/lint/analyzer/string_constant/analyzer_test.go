package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagged(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	Check(p, results)
	testutil.AssertBlocked(t, results, 4)
}

func TestClean(t *testing.T) {
	temporary := testutil.PrepareTestPackage(t, "testdata/src/clean")
	testutil.WriteModFile(t, temporary, "clean")
	p, results := testutil.LoadFromDirectory(t, temporary)
	Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestNoConstant(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/no_constant")
	Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
