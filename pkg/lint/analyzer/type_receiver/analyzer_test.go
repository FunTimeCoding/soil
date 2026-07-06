package type_receiver

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestCheck(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestUnexported(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/unexported")
	Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
