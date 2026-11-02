package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Project() {
	c := gitlab.NewEnvironment()
	defer c.Close()

	for _, p := range c.MustProjects(context.Background()) {
		console.Format("Project: %s\n", p.Format(constant.Format))
	}
}
