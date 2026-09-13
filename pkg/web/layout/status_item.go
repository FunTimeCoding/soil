package layout

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func StatusItem(
	message string,
	kind string,
) gomponents.Node {
	return html.Div(
		html.ID(constant.LayoutStatusLine),
		html.Class(
			join.Empty("container notification-status notification-", kind),
		),
		extended.OutOfBand(constant.SwapOuter),
		gomponents.Attr("role", "status"),
		html.Span(html.Class("notification-message"), gomponents.Text(message)),
	)
}
