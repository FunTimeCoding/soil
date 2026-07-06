package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
)

func Issue() {
	c := sentry.NewEnvironment()
	f := option.Color.Copy()

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
