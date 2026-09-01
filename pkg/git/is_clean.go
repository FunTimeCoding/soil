package git

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/go-git/go-git/v5"
)

func IsClean(
	s git.Status,
	log bool,
) bool {
	for _, status := range s {
		if status.Worktree != git.Unmodified ||
			status.Staging != git.Unmodified {
			if log {
				console.Format(
					"Status: worktree=%b staging=%b\n",
					status.Worktree,
					status.Staging,
				)
			}

			if status.Worktree != git.Untracked &&
				status.Staging != git.Untracked {
				return false
			}
		}
	}

	return true
}
