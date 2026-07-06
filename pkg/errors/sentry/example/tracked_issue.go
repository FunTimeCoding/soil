package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry"
)

func TrackedIssue() {
	c := sentry.NewEnvironment()
	f := option.Color.Copy()

	for _, i := range c.MustTrackedIssues() {
		fmt.Printf("Issue: %s\n", i.Format(f))
	}
}
