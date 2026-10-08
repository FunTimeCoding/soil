package unexported_struct

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unexported_struct"
	"testing"
)

func TestFlagsEveryUnexportedStruct(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	unexported_struct.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedAt(t, results, "plan.go", 3)
	testutil.AssertBlockedAt(t, results, "worker.go", 3)
	testutil.AssertBlockedContains(t, results, "struct plan is unexported")
}

func TestGeneratedStructIsIgnored(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/generated")
	unexported_struct.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestFunctionLocalStructIsIgnored(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/local")
	unexported_struct.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
