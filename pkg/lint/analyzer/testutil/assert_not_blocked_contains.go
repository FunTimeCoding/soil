package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
	"testing"
)

func AssertNotBlockedContains(
	t *testing.T,
	results *output.Results,
	substring string,
) {
	t.Helper()

	for _, c := range results.Entries {
		if !c.Fixed && !c.Planned && strings.Contains(c.Text, substring) {
			t.Errorf("blocked result containing %q: %s", substring, c.Text)
		}
	}
}
