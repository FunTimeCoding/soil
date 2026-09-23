package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

func writeProcfile(
	t *testing.T,
	content string,
) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "Procfile")
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0644))

	return path
}
