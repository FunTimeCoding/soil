package module_graph

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func Excluded(
	root string,
	directory string,
) bool {
	if directory == root {
		return false
	}

	return excludedName(filepath.Base(directory)) ||
		system.NestedModule(root, directory)
}
