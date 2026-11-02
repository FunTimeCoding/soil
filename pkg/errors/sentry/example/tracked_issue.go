package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func TrackedIssue() {
	c := sentry.NewEnvironment()
	defer c.Close()
	f := consoleConstant.ColorFormat.Copy()

	for _, i := range c.MustTrackedIssues() {
		console.Format("Issue: %s\n", i.Format(f))
	}
}
