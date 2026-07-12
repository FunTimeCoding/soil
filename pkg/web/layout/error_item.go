package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ErrorItem(
	message string,
	event string,
) gomponents.Node {
	nodes := []gomponents.Node{
		gomponents.Attr("role", "alert"),
		html.Class("notification-error"),
		html.Span(
			html.Class("notification-message"),
			gomponents.Text(message),
		),
	}

	if event != "" {
		nodes = append(
			nodes,
			html.Small(
				html.Class("notification-event"),
				gomponents.Text(event),
			),
		)
	}

	nodes = append(
		nodes,
		html.Button(
			html.Class("notification-close"),
			gomponents.Attr("onclick", "this.parentElement.remove()"),
			gomponents.Text("×"),
		),
	)

	return html.Div(nodes...)
}
