package index

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (w *Workspace) Reaches(pattern string) bool {
	prefix, subtree := strings.CutSuffix(pattern, constant.RecursivePattern)

	if !subtree {
		_, okay := w.graph.Nodes[pattern]

		return okay
	}

	for path := range w.graph.Nodes {
		if path == prefix || strings.HasPrefix(path, join.Empty(prefix, "/")) {
			return true
		}
	}

	return false
}
