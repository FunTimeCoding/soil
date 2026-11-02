package service

import "path/filepath"

func rootRelative(
	root string,
	path string,
) string {
	if !filepath.IsAbs(path) {
		return path
	}

	result, e := filepath.Rel(root, path)

	if e != nil {
		return path
	}

	return filepath.ToSlash(result)
}
