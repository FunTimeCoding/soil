package page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings"
)

func (p *Page) PrintConsole() {
	fmt.Println(
		console.Cyan(
			"%s",
			strings.PrefixMultiline(bodyToMarkdown(p.Raw.Body), "> "),
		),
	)
}
