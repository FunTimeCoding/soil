package status_dot

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (d *StatusDot) Node() gomponents.Node {
	return html.Span(
		html.Class(constant.StatusDotClass),
		gomponents.If(d.title != "", html.Title(d.title)),
		html.Style(fmt.Sprintf("--dot-color: %s;", d.color)),
	)
}
