package reach

import "path/filepath"

func resolvePath(
	modules map[string]string,
	module string,
	path string,
) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(modules[module], path)
}
