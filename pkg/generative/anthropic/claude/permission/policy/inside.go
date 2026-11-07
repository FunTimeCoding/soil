package policy

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (w *Walk) inside(path string) bool {
	for _, root := range w.roots {
		if strings.HasPrefix(path, join.Empty(root, "/")) {
			return true
		}
	}

	return false
}
