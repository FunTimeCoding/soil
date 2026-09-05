package defer_close

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/defer_close"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	defer_close.Check(p, results)
	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(t, results, "PanicClose")
}
