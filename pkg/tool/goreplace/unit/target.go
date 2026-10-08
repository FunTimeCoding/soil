package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func target(
	t *testing.T,
	content string,
) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "target.md")
	assert.FatalOnError(t, os.WriteFile(path, []byte(content), 0o644))

	return path
}
