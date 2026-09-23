package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func AssertPositioned(
	t *testing.T,
	results *output.Results,
) {
	t.Helper()

	for _, c := range results.Entries {
		if c.Fixed || c.Planned {
			continue
		}

		if c.Type != constant.ConcernLine || c.Line <= 0 {
			t.Errorf("blocked concern without a line: %s: %s", c.Path, c.Text)
		}
	}
}
