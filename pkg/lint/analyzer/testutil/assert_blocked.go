package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func AssertBlocked(
	t *testing.T,
	results *output.Results,
	count int,
) {
	t.Helper()
	blocked := countBlocked(results)

	if blocked != count {
		t.Errorf("expected %d blocked, got %d", count, blocked)

		for _, c := range results.Entries {
			if !c.Fixed {
				t.Logf("  blocked: %s: %s", c.Path, c.Text)
			}
		}
	}
}
