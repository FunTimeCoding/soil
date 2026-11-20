package lint_tester

import (
	"os"
	"path/filepath"
)

func (o *Tester) WriteTestFile(
	name string,
	content string,
) string {
	o.t.Helper()
	path := filepath.Join(o.t.TempDir(), name)
	e := os.WriteFile(path, []byte(content), 0644)

	if e != nil {
		o.t.Fatalf("write: %s", e)
	}

	return path
}
