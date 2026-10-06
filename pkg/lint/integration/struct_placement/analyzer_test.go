package struct_placement

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/struct_placement"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagsEveryStructBesideAReceiverStruct(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/flagged")
	struct_placement.Check(p, results)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedAt(t, results, "payload.go", 3)
	testutil.AssertBlockedAt(t, results, "plan.go", 3)
	testutil.AssertBlockedContains(
		t,
		results,
		"struct Payload shares the package with receiver struct Service",
	)
}

func TestBagWithConstructorsIsClean(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/bag")
	struct_placement.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestPrimitiveReceiverDoesNotTaint(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/primitive")
	struct_placement.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestFunctionLocalStructIsIgnored(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/local")
	struct_placement.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}

func TestGeneratedStructIsIgnored(t *testing.T) {
	p, results := testutil.LoadTestPackage(t, "testdata/src/generated")
	struct_placement.Check(p, results)
	testutil.AssertBlocked(t, results, 0)
}
