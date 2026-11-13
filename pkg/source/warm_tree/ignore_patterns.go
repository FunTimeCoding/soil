package warm_tree

import (
	"github.com/funtimecoding/soil/pkg/git"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"path"
	"slices"
)

func (t *Tree) ignorePatterns(
	directory string,
	domain []string,
	inherited []gitignore.Pattern,
) []gitignore.Pattern {
	result := inherited

	if b, e := t.Read(path.Join(directory, gitConstant.IgnoreFile)); e == nil {
		result = append(result, git.ParsePatterns(b, domain)...)
	}

	for _, name := range t.Children(directory) {
		child := path.Join(directory, name)

		if e := t.entries[child]; !e.Directory || e.Linked {
			continue
		}

		inner := append(slices.Clone(domain), name)

		if gitignore.NewMatcher(result).Match(inner, true) {
			continue
		}

		result = append(result, t.ignorePatterns(child, inner, nil)...)
	}

	return result
}
