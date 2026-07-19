package naming

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	naming.Check(p, results)
	testutil.AssertBlocked(t, results, 57)
}
