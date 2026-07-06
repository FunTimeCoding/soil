package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func addForm() gomponents.Node {
	return html.Form(
		htmx.Post(constant.AddEntryPath),
		htmx.Target("#add-form"),
		htmx.Swap("outerHTML"),
		html.ID("add-form"),
		html.Label(
			gomponents.Text("Action (required)"),
			html.Input(
				html.Type("text"),
				html.Name(constant.Action),
				html.Required(),
				html.Placeholder("e.g. restarted web server"),
			),
		),
		html.Label(
			gomponents.Text("User (required)"),
			html.Input(
				html.Type("text"),
				html.Name(constant.User),
				html.Required(),
				html.Placeholder("e.g. jdoe"),
			),
		),
		html.Div(
			html.Class("grid"),
			html.Label(
				gomponents.Text("System"),
				html.Input(
					html.Type("text"),
					html.Name(constant.System),
					html.Placeholder("e.g. worker1"),
				),
			),
			html.Label(
				gomponents.Text("Service"),
				html.Input(
					html.Type("text"),
					html.Name(constant.Service),
					html.Placeholder("e.g. nginx"),
				),
			),
		),
		html.Label(
			gomponents.Text("Description"),
			html.Textarea(
				html.Name(constant.Description),
				html.Placeholder("Optional details..."),
				html.Rows("3"),
			),
		),
		html.Button(
			html.Type("submit"),
			gomponents.Text(constant.AddEntryTitle),
		),
	)
}
