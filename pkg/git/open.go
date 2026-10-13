package git

import "github.com/go-git/go-git/v5"

func Open(path string) (*git.Repository, error) {
	return git.PlainOpenWithOptions(
		path,
		&git.PlainOpenOptions{DetectDotGit: true, EnableDotGitCommonDir: true},
	)
}
