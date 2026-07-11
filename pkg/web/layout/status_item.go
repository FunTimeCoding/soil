package layout

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func StatusItem(message string, kind string) gomponents.Node {
	return html.Div(
		html.ID(StatusLine),
		html.Class(
			join.Empty(
				"container notification-status notification-",
				kind,
			),
		),
		gomponents.Attr("hx-swap-oob", "outerHTML"),
		gomponents.Attr("role", "status"),
		html.Span(
			html.Class("notification-message"),
			gomponents.Text(message),
		),
	)
}
