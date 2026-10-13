package git

import (
	"github.com/go-git/go-git/v5"
	"path/filepath"
	"sort"
)

func ModifiedFiles(path string) ([]string, error) {
	r, e := Open(path)

	if e != nil {
		return nil, e
	}

	t, f := r.Worktree()

	if f != nil {
		return nil, f
	}

	status, g := t.Status()

	if g != nil {
		return nil, g
	}

	directory, h := filepath.EvalSymlinks(path)

	if h != nil {
		return nil, h
	}

	root, i := filepath.EvalSymlinks(t.Filesystem.Root())

	if i != nil {
		return nil, i
	}

	scope, j := filepath.Rel(root, directory)

	if j != nil {
		return nil, j
	}

	var result []string

	for file, s := range status {
		if s.Worktree == git.Untracked ||
			(s.Worktree == git.Unmodified && s.Staging == git.Unmodified) {
			continue
		}

		if relative, inside := scoped(file, filepath.ToSlash(scope)); inside {
			result = append(result, relative)
		}
	}

	sort.Strings(result)

	return result, nil
}
