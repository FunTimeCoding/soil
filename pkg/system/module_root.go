package system

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"path/filepath"
)

func ModuleRoot(
	directory string,
	limit string,
) string {
	for {
		if FileExists(filepath.Join(directory, constant.ModFile)) {
			return directory
		}

		if filepath.Clean(directory) == filepath.Clean(limit) {
			return limit
		}

		parent := filepath.Dir(directory)

		if parent == directory {
			return limit
		}

		directory = parent
	}
}
