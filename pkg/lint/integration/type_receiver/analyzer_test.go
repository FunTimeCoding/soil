package type_receiver

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/type_receiver"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	type_receiver.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestUnexported(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/unexported")
	type_receiver.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
}

func TestClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/clean")
	type_receiver.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
