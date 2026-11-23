package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func WrapCell(text string) gomponents.Node {
	return html.Td(
		html.Class(constant.LayoutWrapCellClass),
		gomponents.Text(text),
	)
}
