package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func subServices(
	v *virtual_file_system.System,
	directory string,
) []string {
	var result []string

	for _, name := range v.MustReadDirectory(directory) {
		child := filepath.Join(directory, name)

		for _, path := range sourceFiles(v, child) {
			if f := parseWebFile(v, path); f != nil && declaresService(f) {
				result = append(result, name)

				break
			}
		}
	}

	return result
}
