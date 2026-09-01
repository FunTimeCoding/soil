package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Label() {
	c := confluence.NewEnvironment()
	f := constant.ConfluenceDense

	for _, l := range c.MustLabels() {
		console.Format("Label: %+v\n", l)
	}

	c.SetLabels([]string{"favourite"})

	for _, o := range c.MustLabeled() {
		console.Line(o.Format(f))
	}
}
