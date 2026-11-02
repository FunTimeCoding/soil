package revision

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"path"
)

func importPath(
	module string,
	directory string,
) string {
	if directory == soil.CurrentDirectory {
		return module
	}

	return path.Join(module, directory)
}
