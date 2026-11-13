package repository

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path"
	"strings"
)

func (r *Repository) Exists(p string) bool {
	if p == "" {
		return false
	}

	if r.Files.Has(p) {
		return true
	}

	clean := path.Clean(p)

	if r.Directories != nil && !InGitDirectory(clean) {
		if s, found := r.siblingOf(clean); found {
			return r.walkSibling(s).Directories[clean]
		}

		if !strings.HasPrefix(clean, soil.ParentDirectory) &&
			!r.inNestedModule(clean) {
			return r.Directories[clean]
		}
	}

	if r.Files.DirectoryExists(p) || system.DirectoryExists(r.Absolute(p)) {
		return true
	}

	return r.inNestedModule(p) && system.FileExists(r.Absolute(p))
}
