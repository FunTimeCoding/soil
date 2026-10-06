package module_graph

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path"
	"path/filepath"
)

func replacedPath(
	replaced map[string]string,
	directory string,
) (string, bool) {
	for module, root := range replaced {
		relative, e := filepath.Rel(root, directory)

		if e != nil || !filepath.IsLocal(relative) {
			continue
		}

		if relative == constant.CurrentDirectory {
			return module, true
		}

		return path.Join(module, filepath.ToSlash(relative)), true
	}

	return "", false
}
