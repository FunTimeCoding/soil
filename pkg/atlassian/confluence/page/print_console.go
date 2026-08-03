package page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/strings"
)

func (p *Page) PrintConsole() {
	fmt.Println(
		constant.Cyan(
			"%s",
			strings.PrefixMultiline(bodyToMarkdown(p.Raw.Body), "> "),
		),
	)
}
