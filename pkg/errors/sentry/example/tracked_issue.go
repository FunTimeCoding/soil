package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func TrackedIssue() {
	c := sentry.NewEnvironment()
	f := constant.ColorFormat.Copy()

	for _, i := range c.MustTrackedIssues() {
		console.Format("Issue: %s\n", i.Format(f))
	}
}
