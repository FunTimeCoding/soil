package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func sourceFiles(
	v *virtual_file_system.System,
	directory string,
) []string {
	var result []string

	for _, f := range treeSourceFiles(v, directory) {
		if filepath.Dir(f) == directory {
			result = append(result, f)
		}
	}

	return result
}
