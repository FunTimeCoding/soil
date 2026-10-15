package forbidden_call

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	forbidden_call.Check(p, results)
	testutil.AssertBlocked(t, results, 3)
	testutil.AssertBlockedContains(t, results, "pkg/system/run")
}

func TestAMarkerForAnotherKeyDoesNotSuppress(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/example")
	forbidden_call.Check(p, results)
	testutil.AssertBlockedAt(t, results, "other_key.go", 6)
}
