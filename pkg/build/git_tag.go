package build

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/git"
)

func GitTag() string {
	p := git.FindDirectory()
	result := git.LatestTag(p)

	if result == "" {
		errors.Warning("no tag found: %s", p)
	}

	return result
}
