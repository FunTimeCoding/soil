package common

import (
	"fmt"
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
		fmt.Printf("repository not found: %s/%s\n", owner, repository)
		os.Exit(1)
	}

	return p
}
