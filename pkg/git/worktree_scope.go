package git

import (
	"github.com/go-git/go-git/v5"
	"path/filepath"
)

func worktreeScope(
	r *git.Repository,
	path string,
) (string, error) {
	t, e := r.Worktree()

	if e != nil {
		return "", e
	}

	absolute, f := filepath.Abs(path)

	if f != nil {
		return "", f
	}

	directory, g := filepath.EvalSymlinks(absolute)

	if g != nil {
		return "", g
	}

	root, h := filepath.EvalSymlinks(t.Filesystem.Root())

	if h != nil {
		return "", h
	}

	scope, i := filepath.Rel(root, directory)

	if i != nil {
		return "", i
	}

	return filepath.ToSlash(scope), nil
}
