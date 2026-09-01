package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Export() {
	c := confluence.NewEnvironment()
	f := constant.ConfluenceDense

	for _, p := range c.MustChildPages(c.DefaultSpace(), c.DefaultPage()) {
		console.Line(p.Format(f))
		c.Export(p, fmt.Sprintf("fixture/wiki/example/%s.json", p.Name))
	}
}
