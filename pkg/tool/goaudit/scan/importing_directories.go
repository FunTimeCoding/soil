package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
	"slices"
	"strconv"
)

func importingDirectories(
	v *virtual_file_system.System,
	directory string,
	matches func(string) bool,
) []string {
	var result []string

	for _, path := range treeSourceFiles(v, directory) {
		f := parseWebFile(v, path)

		if f == nil {
			continue
		}

		for _, i := range f.Imports {
			value, e := strconv.Unquote(i.Path.Value)

			if e != nil || !matches(value) {
				continue
			}

			if d := filepath.Dir(path); !slices.Contains(result, d) {
				result = append(result, d)
			}
		}
	}

	return result
}
