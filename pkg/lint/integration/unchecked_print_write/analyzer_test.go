package unchecked_print_write

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unchecked_print_write"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	unchecked_print_write.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
}
