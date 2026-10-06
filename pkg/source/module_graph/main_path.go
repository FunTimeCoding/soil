package module_graph

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path"
	"path/filepath"
	"strings"
)

func mainPath(
	root string,
	module string,
	directory string,
) (string, bool) {
	relative, e := filepath.Rel(root, directory)

	if e != nil || !filepath.IsLocal(relative) {
		return "", false
	}

	if relative == constant.CurrentDirectory {
		return module, true
	}

	current := root

	for _, name := range strings.Split(relative, string(filepath.Separator)) {
		current = filepath.Join(current, name)

		if Excluded(root, current) {
			return "", false
		}
	}

	return path.Join(module, filepath.ToSlash(relative)), true
}
