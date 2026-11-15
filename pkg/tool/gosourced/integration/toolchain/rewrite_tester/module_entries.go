package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
	"testing"
)

func ModuleEntries(
	t *testing.T,
	r *output.Results,
) {
	t.Helper()

	for _, c := range r.Entries {
		if filepath.IsAbs(c.Path) {
			t.Errorf("rewrite target outside module: %s", c.Path)
		}
	}
}
