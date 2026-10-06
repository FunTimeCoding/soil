package module_graph

import (
	rootConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
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
		strings.HasPrefix(name, rootConstant.CurrentDirectory) ||
		strings.HasPrefix(name, "_") ||
		system.FileExists(filepath.Join(directory, constant.ModFile))
}
