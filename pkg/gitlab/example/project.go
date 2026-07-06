package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Project() {
	for _, p := range gitlab.NewEnvironment().MustProjects() {
		fmt.Printf("Project: %s\n", p.Format(constant.Format))
	}
}
