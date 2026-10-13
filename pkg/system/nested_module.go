package system

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"path/filepath"
)

func NestedModule(
	root string,
	directory string,
) bool {
	if filepath.Clean(directory) == filepath.Clean(root) {
		return false
	}

	return FileExists(filepath.Join(directory, constant.ModFile))
}
