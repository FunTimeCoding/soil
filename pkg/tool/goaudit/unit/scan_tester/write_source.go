package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func WriteSource(
	base string,
	path string,
) {
	system.MakeDirectory(filepath.Join(base, filepath.Dir(path)))
	system.WriteFile(filepath.Join(base, path), []byte("package alfa\n"), 0644)
}
