package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func Tree(path string) *git.Worktree {
	t, e := MustOpen(path).Worktree()
	errors.PanicOnError(e)

	return t
}
