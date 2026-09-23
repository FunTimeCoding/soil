package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
	"testing"
)

func assertEntriesInsideModule(
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
