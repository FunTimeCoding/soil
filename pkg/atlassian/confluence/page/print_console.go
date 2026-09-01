package page

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/strings"
)

func (p *Page) PrintConsole() {
	console.Line(
		constant.Cyan(
			"%s",
			strings.PrefixMultiline(bodyToMarkdown(p.Raw.Body), "> "),
		),
	)
}
