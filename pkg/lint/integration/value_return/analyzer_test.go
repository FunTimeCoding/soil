package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	testutil.AssertBlocked(t, checked(t), 5)
}

func TestMarkedTypeIsNotFlagged(t *testing.T) {
	testutil.AssertNotBlockedContains(t, checked(t), "MarkedByValue")
}

func TestDirectiveWithoutReasonIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(t, checked(t), "BareDirectiveByValue")
}

func TestTypeDeclaredBesideMarkedOneIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(t, checked(t), "UnmarkedByValue")
}
