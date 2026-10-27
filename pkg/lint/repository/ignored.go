package repository

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func (r *Repository) Ignored(path string) bool {
	clean := filepath.Clean(path)

	for sibling, ignored := range r.siblingIgnored {
		inner, found := strings.CutPrefix(
			clean,
			join.Empty(sibling, constant.Slash),
		)

		if found {
			return ignored(inner)
		}
	}

	return r.ignored(clean)
}
