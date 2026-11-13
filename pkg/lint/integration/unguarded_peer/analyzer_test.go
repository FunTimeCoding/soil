package unguarded_peer

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestUnguardedPeerFlagsCallsOutsideGuard(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlocked(t, results, 6)

	for _, line := range []int{9, 16, 23, 45, 68, 77} {
		testutil.AssertLineBlocked(t, results, "flagged.go", line)
	}

	testutil.AssertBlockedContains(t, results, "c.Book reaches a guarded peer")
}
