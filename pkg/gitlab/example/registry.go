package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Registry() {
	c := gitlab.NewEnvironment()
	r := c.MustRegistryRepositories(
		c.MustProjectByName(
			environment.Required(constant.GroupEnvironment),
			environment.Required(constant.ProjectEnvironment),
		).Identifier,
		false,
	)
	fmt.Printf("Repositories: %d\n", len(r))

	for _, r := range r {
		fmt.Printf("Registry: %d %s\n", r.ID, r.Name)
	}
}
