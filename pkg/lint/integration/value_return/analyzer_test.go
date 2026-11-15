package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestBlocked(t *testing.T) {
	testutil.AssertBlocked(t, lint_tester.ValueReturn(t), 5)
}

func TestMarkedTypeIsNotFlagged(t *testing.T) {
	testutil.AssertNotBlockedContains(
		t,
		lint_tester.ValueReturn(t),
		"MarkedByValue",
	)
}

func TestDirectiveWithoutReasonIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(
		t,
		lint_tester.ValueReturn(t),
		"BareDirectiveByValue",
	)
}

func TestTypeDeclaredBesideMarkedOneIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(
		t,
		lint_tester.ValueReturn(t),
		"UnmarkedByValue",
	)
}
