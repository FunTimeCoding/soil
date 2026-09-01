package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Organization() {
	c := sentry.NewEnvironment()

	for _, o := range c.MustOrganizations() {
		console.Format("Organization: %+v\n", o)

		for _, t := range c.MustTeams(o.Slug) {
			console.Format("Team: %+v\n", t)
		}
	}
}
