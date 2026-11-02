package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func Issue() {
	g := gitlab.NewEnvironment()
	defer g.Close()
	x := context.Background()

	for _, i := range g.MustIssues(x) {
		p := g.MustProject(x, i.Project)
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
