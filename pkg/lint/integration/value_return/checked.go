package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/value_return"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func checked(t *testing.T) *output.Results {
	t.Helper()
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	value_return.Check(p, results)

	return results
}
