package repository

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/sibling_walk"
	"path"
	"strings"
)

func (r *Repository) walkSibling(sibling string) *sibling_walk.Walk {
	if cached, found := r.siblingWalks[sibling]; found {
		return cached
	}

	result := sibling_walk.New()
	t := r.trees.Tree(r.Absolute(sibling))
	result.Directories[path.Clean(sibling)] = true

	for _, p := range t.Paths() {
		e, _ := t.Stat(p)
		relative := path.Join(sibling, p)

		if e.Directory {
			result.Directories[relative] = true
		} else if !strings.HasSuffix(p, soil.MarkdownExtension) {
			result.Files = append(result.Files, relative)
		}
	}

	r.siblingWalks[sibling] = result

	return result
}
