package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func write(
	root string,
	name string,
	content string,
) {
	path := filepath.Join(root, name)
	system.MakeDirectory(filepath.Dir(path))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o600))
}
