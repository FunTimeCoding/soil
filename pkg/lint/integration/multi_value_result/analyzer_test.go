package multi_value_result

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/multi_value_result"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	multi_value_result.Check(p, results)
	testutil.AssertBlocked(t, results, 7)
}
