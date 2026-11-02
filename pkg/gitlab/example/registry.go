package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Registry() {
	a := argument.NewSimple("gitlab-registry")
	a.String(argumentConstant.Owner, "", "group or user")
	a.String(argumentConstant.Repository, "", "repository name")
	a.ParseSimple()
	owner := a.GetString(argumentConstant.Owner)
	repository := a.GetString(argumentConstant.Repository)
	c := gitlab.NewEnvironment()
	defer c.Close()
	x := context.Background()
	r := c.MustRegistryRepositories(
		x,
		c.MustProjectByName(x, owner, repository).Identifier,
		false,
	)
	console.Format("Repositories: %d\n", len(r))

	for _, r := range r {
		console.Format("Registry: %d %s\n", r.Identifier, r.Name)
	}
}
