package reach

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"path/filepath"
	"strings"
)

func rootRelative(
	root string,
	directory string,
) (string, bool) {
	relative, e := filepath.Rel(root, directory)

	if e != nil || strings.HasPrefix(relative, soil.ParentDirectory) {
		return "", false
	}

	for d := directory; d != root && strings.HasPrefix(d, root); d = filepath.Dir(d) {
		if module_graph.Excluded(root, d) {
			return "", false
		}
	}

	return relative, true
}
