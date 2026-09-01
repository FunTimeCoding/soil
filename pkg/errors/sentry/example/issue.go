package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Issue() {
	c := sentry.NewEnvironment()
	f := consoleConstant.ColorFormat.Copy()

	for _, i := range c.MustIssuesSimple(true) {
		console.Format("Issue: %s\n", i.Format(f))
	}

	if false {
		for _, o := range c.MustOrganizations() {
			console.Format("Organization: %s\n", o.Name)

			for _, p := range c.MustOrganizationProjects(o.Slug) {
				console.Format("Project: %s\n", p.Name)

				for _, i := range c.MustIssues(
					o.Slug,
					p.Identifier,
					constant.PeriodFortnight,
				) {
					console.Format("Issue: %s\n", i.Format(f))
				}
			}
		}
	}
}
