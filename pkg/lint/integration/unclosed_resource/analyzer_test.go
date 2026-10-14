package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestFlagged(t *testing.T) {
	results := checkModule(t, "testdata/src/flagged")
	testutil.AssertBlocked(t, results, 10)
	testutil.AssertBlockedContains(t, results, "PanicClose(r.Body)")
	testutil.AssertBlockedContains(t, results, "local is never closed")
	testutil.AssertBlockedContains(t, results, "c is never closed")
}

func TestClean(t *testing.T) {
	testutil.AssertBlocked(t, checkModule(t, "testdata/src/clean"), 0)
}

func TestArrangedInSeparateLoad(t *testing.T) {
	testutil.AssertBlocked(t, checkModule(t, "testdata/src/elsewhere"), 0)
}

func TestArrangedByConstructor(t *testing.T) {
	testutil.AssertBlocked(t, checkModule(t, "testdata/src/arranged"), 0)
}

func TestTransferredWithoutRelease(t *testing.T) {
	results := checkModule(t, "testdata/src/transferred")
	testutil.AssertBlocked(t, results, 13)
	testutil.AssertBlockedContains(
		t,
		results,
		"Bag holds resource but has no Close to release it",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Pair.Close never releases second",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Hider.reader holds *os.File behind io.Reader, which has no Close",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Pool.Close never releases members",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"List.Close never releases items",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Watched.Close never releases resource",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Shell.Close never releases inner",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Outer.Close is promoted from Base and never releases extra",
	)
}

func TestTransferredAndReleased(t *testing.T) {
	testutil.AssertBlocked(t, checkModule(t, "testdata/src/released"), 0)
}

func TestTransferredAcrossPackages(t *testing.T) {
	results := checkModule(t, "testdata/src/handed")
	testutil.AssertBlocked(t, results, 1)
	testutil.AssertBlockedContains(
		t,
		results,
		"Leaky.Close never releases Resource",
	)
}
