package unread_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestUnreadReceiverFlagged(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlocked(t, results, 4)

	for _, name := range []string{"Weight", "Height", "Width", "Shadow"} {
		testutil.AssertBlockedContains(
			t,
			results,
			fmt.Sprintf("%s never reads its receiver", name),
		)
	}
}

func TestUnreadReceiverLeavesReadersAndImplementations(t *testing.T) {
	results := analyze(t)

	for _, name := range []string{"Name ", "Describe ", "Error ", "Length "} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}
