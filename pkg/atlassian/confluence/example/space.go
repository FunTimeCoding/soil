package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Space() {
	f := constant.ConfluenceFormat

	for _, s := range confluence.NewEnvironment().MustSpaces() {
		console.Line(s.Format(f))
	}
}
