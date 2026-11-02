package common

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"os"
)

func FindProjectOrExit(
	x context.Context,
	c face.GitLabClient,
	owner string,
	repository string,
) *project.Project {
	p := c.MustProjectByName(x, owner, repository)

	if p == nil {
		console.Format("repository not found: %s/%s\n", owner, repository)
		os.Exit(1)
	}

	return p
}
