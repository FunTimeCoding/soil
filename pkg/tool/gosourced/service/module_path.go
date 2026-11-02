package service

import "path/filepath"

func modulePath(
	directory string,
	path string,
) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}

	return filepath.Join(directory, path)
}
