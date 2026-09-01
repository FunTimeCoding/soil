package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Project() {
	for _, p := range gitlab.NewEnvironment().MustProjects() {
		console.Format("Project: %s\n", p.Format(constant.Format))
	}
}
