package unread_receiver

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unread_receiver"
	"testing"
)

func TestUnreadReceiverFlagged(t *testing.T) {
	results := testutil.AnalyzeFaces(
		t,
		testutil.PrepareTestPackage(t, "testdata/src/example"),
		unread_receiver.Check,
	)
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
	results := testutil.AnalyzeFaces(
		t,
		testutil.PrepareTestPackage(t, "testdata/src/example"),
		unread_receiver.Check,
	)

	for _, name := range []string{"Name ", "Describe ", "Error ", "Length "} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}
