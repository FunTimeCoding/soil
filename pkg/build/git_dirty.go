package build

import "github.com/funtimecoding/soil/pkg/git"

func GitDirty() bool {
	return !git.IsClean(git.Status(git.FindDirectory()), false)
}
