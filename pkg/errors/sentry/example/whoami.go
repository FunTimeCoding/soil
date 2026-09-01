package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func Whoami() {
	c := sentry.NewEnvironment()
	u := c.MustWhoami()
	console.Format("User: %+v\n", u)
}
