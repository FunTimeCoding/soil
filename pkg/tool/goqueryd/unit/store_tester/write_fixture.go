package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func (o *Tester) WriteFixture(
	directory string,
	name string,
	content string,
) {
	o.t.Helper()
	path := filepath.Join(directory, name)
	errors.PanicOnError(os.MkdirAll(filepath.Dir(path), 0o755))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o644))
}
