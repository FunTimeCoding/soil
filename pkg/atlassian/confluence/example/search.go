package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Search() {
	c := confluence.NewEnvironment()
	f := constant.ConfluenceDense

	if true {
		for _, r := range c.MustSearch("space=%s", c.DefaultSpace()) {
			console.Line(r.Format(f))
		}
	}

	if false {
		// Working syntax examples
		c.MustSearch("favorite=currentUser()")
		c.MustSearch(`label IN ("ExampleLabel")`)
		c.MustSearch("creator IN (currentUser())")
		c.MustSearch("creator=currentUser()")
		c.MustSearch("creator=currentUser()")
		c.MustSearch("space=EXAMPLE")
		c.MustSearch("space=EXAMPLE")
		c.MustSearch("watcher=currentUser()")
	}
}
