package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Issue() {
	g := gitlab.NewEnvironment()

	for _, i := range g.MustIssues() {
		p := g.MustProject(i.Project)
		console.Format("Project: %s\n", p.Format(constant.Format))
		console.Format("  Issue: %s\n", i.Format(constant.Format))
		console.Format("  %s\n", i.Raw.WebURL)

		if false {
			console.Format(
				"  %s\n",
				consoleConstant.Magenta("%s", i.Raw.Description),
			)
		}
	}
}
