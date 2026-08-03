package example

import (
	"fmt"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Issue() {
	c := sentry.NewEnvironment()
	f := console.ColorFormat.Copy()

	for _, i := range c.MustIssuesSimple(true) {
		fmt.Printf("Issue: %s\n", i.Format(f))
	}

	if false {
		for _, o := range c.MustOrganizations() {
			fmt.Printf("Organization: %s\n", o.Name)

			for _, p := range c.MustOrganizationProjects(o.Slug) {
				fmt.Printf("Project: %s\n", p.Name)

				for _, i := range c.MustIssues(
					o.Slug,
					p.ID,
					constant.PeriodFortnight,
				) {
					fmt.Printf("Issue: %s\n", i.Format(f))
				}
			}
		}
	}
}
