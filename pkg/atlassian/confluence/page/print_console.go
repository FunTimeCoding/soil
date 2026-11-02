package page

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/strings"
)

func (p *Page) PrintConsole() {
	console.Line(
		consoleConstant.Cyan(
			"%s",
			strings.PrefixMultiline(bodyToMarkdown(p.Raw.Body), "> "),
		),
	)
}
