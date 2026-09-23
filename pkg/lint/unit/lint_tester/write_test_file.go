package lint_tester

import (
	"os"
	"path/filepath"
	"testing"
)

func WriteTestFile(
	t *testing.T,
	name string,
	content string,
) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	e := os.WriteFile(path, []byte(content), 0644)

	if e != nil {
		t.Fatalf("write: %s", e)
	}

	return path
}
