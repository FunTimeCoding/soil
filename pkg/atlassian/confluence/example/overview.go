package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Overview() {
	c := confluence.NewEnvironment(confluence.WithVerbose(true))
	f := constant.ConfluenceFormat
	console.Line("Space")

	for _, s := range c.MustSpaces() {
		console.Line(s.Format(f))
	}

	console.Line("Page")

	for _, a := range c.MustPagesBySpaceName(c.DefaultSpace()) {
		console.Line(a.Format(f))
	}
}
