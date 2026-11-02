package forwarding_function

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forwarding_function"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestForwardingFunction(t *testing.T) {
	results := analyze(t, forwarding_function.Check)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedContains(t, results, "small only forwards to NewSize")
	testutil.AssertBlockedContains(t, results, "plain only forwards to add")
}

func TestDaemonForwardingToServiceFlagged(t *testing.T) {
	results := analyze(t, forwarding_function.CheckDaemon)
	testutil.AssertBlocked(t, results, 2)
	testutil.AssertBlockedContains(
		t,
		results,
		"Renew only forwards to the loan service - hand the surface the service",
	)
	testutil.AssertBlockedContains(
		t,
		results,
		"Packages only forwards to the service - hand the surface the service",
	)
}
