package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Project() {
	for _, p := range sentry.NewEnvironment().MustProjects() {
		console.Format("%+v\n", p)
	}
}
