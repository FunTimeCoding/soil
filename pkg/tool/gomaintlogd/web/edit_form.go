package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func editForm(e *entry.Entry, message string) gomponents.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)
	var alert gomponents.Node

	if message != "" {
		alert = layout.Alert(message)
	}

	return html.Tr(
		html.ID(fmt.Sprintf("detail-%d", e.ID)),
		html.Class("detail-row"),
		html.Td(
			gomponents.Attr("colspan", "6"),
			alert,
			html.Form(
				html.Class("edit-form"),
				htmx.Post(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
				htmx.Target(target),
				htmx.Swap("outerHTML"),
				html.Div(
					html.Class("grid"),
					html.Label(
						gomponents.Text("Action"),
						html.Input(
							html.Type("text"),
							html.Name(constant.Action),
							html.Value(e.Action),
							html.Required(),
						),
					),
					html.Label(
						gomponents.Text("User"),
						html.Input(
							html.Type("text"),
							html.Name(constant.User),
							html.Value(e.User),
							html.Required(),
						),
					),
					html.Label(
						gomponents.Text("Timestamp"),
						html.Input(
							html.Type("datetime-local"),
							html.Name(constant.Timestamp),
							html.Value(
								e.Timestamp.Format("2006-01-02T15:04"),
							),
						),
					),
				),
				html.Div(
					html.Class("grid"),
					html.Label(
						gomponents.Text("System"),
						html.Input(
							html.Type("text"),
							html.Name(constant.System),
							html.Value(e.System),
						),
					),
					html.Label(
						gomponents.Text("Service"),
						html.Input(
							html.Type("text"),
							html.Name(constant.Service),
							html.Value(e.Service),
						),
					),
				),
				html.Label(
					gomponents.Text("Description"),
					html.Textarea(
						html.Name(constant.Description),
						html.Rows("4"),
						gomponents.Text(e.Description),
					),
				),
				html.Div(
					html.Class("detail-actions"),
					html.Button(html.Type("submit"), gomponents.Text("Save")),
					html.Button(
						html.Type("button"),
						html.Class("outline secondary"),
						htmx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
						htmx.Target(target),
						htmx.Swap("outerHTML"),
						gomponents.Text("Cancel"),
					),
				),
			),
		),
	)
}
