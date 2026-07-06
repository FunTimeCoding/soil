package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Issue() {
	g := gitlab.NewEnvironment()

	for _, i := range g.MustIssues() {
		p := g.MustProject(i.Project)
		fmt.Printf("Project: %s\n", p.Format(constant.Format))
		fmt.Printf("  Issue: %s\n", i.Format(constant.Format))
		fmt.Printf("  %s\n", i.Raw.WebURL)

		if false {
			fmt.Printf(
				"  %s\n",
				console.Magenta("%s", i.Raw.Description),
			)
		}
	}
}
