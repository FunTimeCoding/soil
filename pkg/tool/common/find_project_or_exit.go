package common

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"os"
)

func FindProjectOrExit(
	c face.Forge,
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
