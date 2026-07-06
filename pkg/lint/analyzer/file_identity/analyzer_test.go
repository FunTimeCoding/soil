package file_identity

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestMultipleIdentities(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/multi")
	Check(p, results)
	testutil.AssertBlocked(t, results, 4)
}

func TestMismatch(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/mismatch")
	Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestMethods(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/methods")
	Check(p, results)
	testutil.AssertBlocked(t, results, 2)
}

func TestInterfaceMethod(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/iface")
	Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
