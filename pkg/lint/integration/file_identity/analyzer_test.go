package file_identity

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/file_identity"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	file_identity.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestMultipleIdentities(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/multi")
	file_identity.Check(p, results)
	testutil.AssertBlocked(t, results, 4)
}

func TestMismatch(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/mismatch")
	file_identity.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestMethods(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/methods")
	file_identity.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
}

func TestInterfaceMethod(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/iface")
	file_identity.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
