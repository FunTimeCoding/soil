package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Organization() {
	c := sentry.NewEnvironment()

	for _, o := range c.MustOrganizations() {
		fmt.Printf("Organization: %+v\n", o)

		for _, t := range c.MustTeams(o.Slug) {
			fmt.Printf("Team: %+v\n", t)
		}
	}
}
