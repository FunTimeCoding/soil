package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
)

func Search() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	if true {
		for _, r := range c.MustSearch("space=%s", c.DefaultSpace()) {
			fmt.Println(r.Format(f))
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
