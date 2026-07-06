package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
	"testing"
)

func AssertBlockedContains(
	t *testing.T,
	results *output.Results,
	substring string,
) {
	t.Helper()

	for _, c := range results.Entries {
		if !c.Fixed && strings.Contains(c.Text, substring) {
			return
		}
	}

	t.Errorf("no blocked result containing %q", substring)

	for _, c := range results.Entries {
		if !c.Fixed {
			t.Logf("  blocked: %s: %s", c.Path, c.Text)
		}
	}
}
