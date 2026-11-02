package reach

import "path/filepath"

func absolute(
	directory string,
	path string,
) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}

	return filepath.Join(directory, path)
}
