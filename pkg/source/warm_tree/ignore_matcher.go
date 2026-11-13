package warm_tree

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"os"
	"path/filepath"
)

func (t *Tree) IgnoreMatcher() func(string) bool {
	var excluded []gitignore.Pattern

	if b, e := os.ReadFile(
		filepath.Join(t.root, filepath.FromSlash(gitConstant.ExcludeFile)),
	); e == nil {
		excluded = git.ParsePatterns(b, nil)
	}

	return git.PatternMatcher(
		t.ignorePatterns(soil.CurrentDirectory, nil, excluded),
	)
}
