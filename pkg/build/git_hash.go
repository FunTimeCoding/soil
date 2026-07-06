package build

import "github.com/funtimecoding/soil/pkg/git"

func GitHash() string {
	return git.ShortHash(git.FindDirectory())
}
