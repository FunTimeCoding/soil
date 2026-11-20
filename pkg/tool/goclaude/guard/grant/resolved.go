package grant

import "path/filepath"

func resolved(path string) string {
	result, e := filepath.EvalSymlinks(path)

	if e != nil {
		return filepath.Clean(path)
	}

	return result
}
