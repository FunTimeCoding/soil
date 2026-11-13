package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ClampCell(text string) gomponents.Node {
	return html.Td(
		html.Class(constant.LayoutClampCellClass),
		html.Span(
			gomponents.Attr(constant.TooltipAttribute, text),
			gomponents.Text(text),
		),
	)
}
