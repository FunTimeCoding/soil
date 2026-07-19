package call_format

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/call_format"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	call_format.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
}
