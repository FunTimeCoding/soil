package system

import "path/filepath"

func InsideDirectory(
	directory string,
	path string,
) bool {
	relative, e := filepath.Rel(directory, path)

	return e == nil && filepath.IsLocal(relative)
}
