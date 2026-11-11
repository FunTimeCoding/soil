package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
	"testing"
)

func assertFixedContains(
	t *testing.T,
	substring string,
	results *output.Results,
) {
	t.Helper()

	for _, c := range results.Entries {
		if c.Fixed && strings.Contains(c.Text, substring) {
			return
		}
	}

	t.Errorf("no fixed result containing %q", substring)

	for _, c := range results.Entries {
		t.Logf("  %s: %s", c.Path, c.Text)
	}
}
