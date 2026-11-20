package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unguarded_peer"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestUnguardedPeerFlagsCallsOutsideGuard(t *testing.T) {
	results := testutil.Analyze(
		t,
		lint_tester.New(t).StubbedSoil(),
		unguarded_peer.Check,
	)
	testutil.AssertBlocked(t, results, 6)

	for _, line := range []int{9, 16, 23, 45, 68, 77} {
		testutil.AssertLineBlocked(t, results, "flagged.go", line)
	}

	testutil.AssertBlockedContains(t, results, "c.Book reaches a guarded peer")
}
