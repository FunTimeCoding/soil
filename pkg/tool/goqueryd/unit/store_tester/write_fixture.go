package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

func WriteFixture(
	t *testing.T,
	directory string,
	name string,
	content string,
) {
	t.Helper()
	path := filepath.Join(directory, name)
	errors.PanicOnError(os.MkdirAll(filepath.Dir(path), 0o755))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o644))
}
