package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"path/filepath"
)

func (o *Tester) ModuleEntries(r *output.Results) {
	o.t.Helper()

	for _, c := range r.Entries {
		if filepath.IsAbs(c.Path) {
			o.t.Errorf("rewrite target outside module: %s", c.Path)
		}
	}
}
