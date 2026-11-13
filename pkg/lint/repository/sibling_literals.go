package repository

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func (r *Repository) siblingLiterals(directory string) []string {
	clean := path.Clean(directory)
	s, found := r.siblingOf(clean)

	if !found || clean == s || InGitDirectory(clean) {
		return siblingContents(r.Absolute(directory))
	}

	prefix := join.Empty(clean, stringsConstant.Slash)
	var result []string

	for _, f := range r.walkSibling(s).Files {
		if f == clean || strings.HasPrefix(f, prefix) {
			result = append(result, r.siblingFile(s, f))
		}
	}

	return result
}
