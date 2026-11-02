package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Space() {
	f := constant.ConfluenceFormat
	c := confluence.NewEnvironment()
	defer c.Close()

	for _, s := range c.MustSpaces() {
		console.Line(s.Format(f))
	}
}
