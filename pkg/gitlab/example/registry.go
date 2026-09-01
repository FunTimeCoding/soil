package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Registry() {
	a := argument.NewSimple("gitlab-registry")
	a.String(constant.Owner, "", "group or user")
	a.String(constant.Repository, "", "repository name")
	a.ParseSimple()
	owner := a.GetString(constant.Owner)
	repository := a.GetString(constant.Repository)
	c := gitlab.NewEnvironment()
	r := c.MustRegistryRepositories(
		c.MustProjectByName(owner, repository).Identifier,
		false,
	)
	console.Format("Repositories: %d\n", len(r))

	for _, r := range r {
		console.Format("Registry: %d %s\n", r.ID, r.Name)
	}
}
