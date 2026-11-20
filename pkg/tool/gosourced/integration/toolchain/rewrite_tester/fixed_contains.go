package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"strings"
)

func (o *Tester) FixedContains(
	substring string,
	results *output.Results,
) {
	o.t.Helper()

	for _, c := range results.Entries {
		if c.Fixed && strings.Contains(c.Text, substring) {
			return
		}
	}

	o.t.Errorf("no fixed result containing %q", substring)

	for _, c := range results.Entries {
		o.t.Logf("  %s: %s", c.Path, c.Text)
	}
}
