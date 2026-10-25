package nil_dependency

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/nil_dependency"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	nil_dependency.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}
