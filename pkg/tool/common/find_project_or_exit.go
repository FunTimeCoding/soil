package common

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"os"
)

func FindProjectOrExit(
	c *gitlab.Client,
	owner string,
	repository string,
) *project.Project {
	p := c.MustProjectByName(owner, repository)

	if p == nil {
		console.Format("repository not found: %s/%s\n", owner, repository)
		os.Exit(1)
	}

	return p
}
