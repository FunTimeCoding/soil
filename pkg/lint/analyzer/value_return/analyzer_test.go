package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestCheck(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}
