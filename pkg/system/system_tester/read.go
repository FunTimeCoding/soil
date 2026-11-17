package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func Read(
	t *testing.T,
	base string,
	path string,
) string {
	t.Helper()
	b, e := os.ReadFile(filepath.Join(base, path))
	assert.FatalOnError(t, e)

	return string(b)
}
