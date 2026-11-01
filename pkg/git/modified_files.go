package git

import (
	"github.com/go-git/go-git/v5"
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

	scope, h := worktreeScope(r, path)

	if h != nil {
		return nil, h
	}

	var result []string

	for file, s := range status {
		if s.Worktree == git.Untracked ||
			(s.Worktree == git.Unmodified && s.Staging == git.Unmodified) {
			continue
		}

		if relative, inside := scoped(file, scope); inside {
			result = append(result, relative)
		}
	}

	sort.Strings(result)

	return result, nil
}
