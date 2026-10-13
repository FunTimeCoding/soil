package repository

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Repository) inNestedModule(path string) bool {
	for _, m := range r.NestedModules {
		if path == m || strings.HasPrefix(path, join.Empty(m, stringsConstant.Slash)) {
			return true
		}
	}

	return false
}
