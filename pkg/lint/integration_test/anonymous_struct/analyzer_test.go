package anonymous_struct

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/anonymous_struct"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	anonymous_struct.Check(p, results)
	testutil.AssertBlocked(t, results, 4)
}
