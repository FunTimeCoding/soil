package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestBlocked(t *testing.T) {
	testutil.AssertBlocked(t, checked(t), 5)
}

func TestAMarkedTypeIsNotFlagged(t *testing.T) {
	testutil.AssertNotBlockedContains(t, checked(t), "MarkedByValue")
}

func TestADirectiveWithoutAReasonIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(t, checked(t), "BareDirectiveByValue")
}

func TestATypeDeclaredBesideAMarkedOneIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(t, checked(t), "UnmarkedByValue")
}
