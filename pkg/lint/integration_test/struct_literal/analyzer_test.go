package struct_literal

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	struct_literal.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
}
