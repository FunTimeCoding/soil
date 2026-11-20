package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	"testing"
)

func TestBlocked(t *testing.T) {
	testutil.AssertBlocked(t, lint_tester.New(t).ValueReturn(), 5)
}

func TestMarkedTypeIsNotFlagged(t *testing.T) {
	testutil.AssertNotBlockedContains(
		t,
		lint_tester.New(t).ValueReturn(),
		"MarkedByValue",
	)
}

func TestDirectiveWithoutReasonIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(
		t,
		lint_tester.New(t).ValueReturn(),
		"BareDirectiveByValue",
	)
}

func TestTypeDeclaredBesideMarkedOneIsStillFlagged(t *testing.T) {
	testutil.AssertBlockedContains(
		t,
		lint_tester.New(t).ValueReturn(),
		"UnmarkedByValue",
	)
}
