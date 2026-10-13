package module_graph

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"strings"
)

func Excluded(
	root string,
	directory string,
) bool {
	if directory == root {
		return false
	}

	name := filepath.Base(directory)

	return name == "testdata" ||
		name == "vendor" ||
		strings.HasPrefix(name, soil.CurrentDirectory) ||
		strings.HasPrefix(name, "_") ||
		system.NestedModule(root, directory)
}
