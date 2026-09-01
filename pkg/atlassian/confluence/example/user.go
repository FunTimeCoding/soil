package example

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func User() {
	console.Line(
		confluence.NewEnvironment().MustUser().Format(constant.ConfluenceFormat),
	)
}
