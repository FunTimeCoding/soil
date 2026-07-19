package string_concatenation

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	string_concatenation.Check(p, results)
	testutil.AssertBlocked(t, results, 5)
}
