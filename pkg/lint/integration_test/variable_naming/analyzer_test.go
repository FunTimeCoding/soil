package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/variable_naming"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	variable_naming.Check(p, results)
	testutil.AssertBlocked(t, results, 38)
}
