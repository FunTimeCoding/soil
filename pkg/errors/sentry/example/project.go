package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Project() {
	c := sentry.NewEnvironment()
	defer c.Close()

	for _, p := range c.MustProjects() {
		console.Format("%+v\n", p)
	}
}
