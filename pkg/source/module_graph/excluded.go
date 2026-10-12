package module_graph

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
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
		system.FileExists(filepath.Join(directory, go_modConstant.ModFile))
}
