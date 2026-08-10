package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func connectionDot() gomponents.Node {
	return html.Span(
		html.ID(constant.LayoutConnection),
		html.Class(constant.LayoutConnectionDisconnected),
	)
}
