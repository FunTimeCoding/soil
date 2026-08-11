package web

import (
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func entryFields(e *entry.Entry) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.Div(
				html.Class("grid"),
				html.Div(
					html.Strong(gomponents.Text("Timestamp: ")),
					gomponents.Text(time.FormatCompact(e.Timestamp)),
				),
				html.Div(
					html.Strong(gomponents.Text("Action: ")),
					gomponents.Text(e.Action),
				),
				html.Div(
					html.Strong(gomponents.Text("User: ")),
					gomponents.Text(e.User),
				),
			),
			html.Div(
				html.Class("grid"),
				html.Div(
					html.Strong(gomponents.Text("System: ")),
					gomponents.Text(e.System),
				),
				html.Div(
					html.Strong(gomponents.Text("Service: ")),
					gomponents.Text(e.Service),
				),
				html.Div(),
			),
			html.Div(
				html.Strong(gomponents.Text("Description:")),
				html.Pre(gomponents.Text(e.Description)),
			),
		},
	)
}
