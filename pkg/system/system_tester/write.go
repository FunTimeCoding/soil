package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
	"testing"
)

func Write(
	t *testing.T,
	base string,
	path string,
	content string,
) {
	t.Helper()
	full := filepath.Join(base, path)
	assert.FatalOnError(t, os.MkdirAll(filepath.Dir(full), 0755))
	assert.FatalOnError(t, os.WriteFile(full, []byte(content), 0644))
}
