package repository_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func (r *Tester) Write(
	name string,
	content string,
) {
	path := filepath.Join(r.Clone, name)
	system.MakeDirectory(filepath.Dir(path))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o600))
}
