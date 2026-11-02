package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func User() {
	c := confluence.NewEnvironment()
	defer c.Close()
	console.Line(c.MustUser().Format(constant.ConfluenceFormat))
}
