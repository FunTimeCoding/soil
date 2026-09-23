package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
	"testing"
)

func AssertBlockedAt(
	t *testing.T,
	results *output.Results,
	file string,
	line int,
) {
	t.Helper()

	for _, c := range results.Entries {
		if !c.Fixed && !c.Planned &&
			filepath.Base(c.Path) == file && c.Line == line {
			return
		}
	}

	t.Errorf("no blocked result at %s:%d", file, line)

	for _, c := range results.Entries {
		if !c.Fixed && !c.Planned {
			t.Logf("  blocked: %s:%d: %s", c.Path, c.Line, c.Text)
		}
	}
}
