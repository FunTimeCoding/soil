package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/strings/upper"
)

func Page() {
	c := confluence.NewEnvironment()
	s := c.DefaultSpace()
	p := c.DefaultPage()

	if false {
		a := c.MustPageBySpaceAndName(s, upper.Charlie)
		c.MustDelete(a.Identifier)
		c.MustImport(s, p, "fixture/wiki/example/", "Charlie.json")
	}
}
