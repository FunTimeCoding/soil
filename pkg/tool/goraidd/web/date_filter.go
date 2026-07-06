package web

import (
	"github.com/funtimecoding/soil/pkg/constant"
	goraidd "github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func dateFilter(startValue, endValue string) gomponents.Node {
	return html.Form(
		html.Class("filter-form"),
		html.Method(constant.Get),
		html.Action(goraidd.LogsPath),
		html.Div(
			html.Class("grid"),
			html.Label(
				gomponents.Text("From"),
				html.Input(
					html.Type("datetime-local"),
					html.Name("start"),
					html.Value(startValue),
				),
			),
			html.Label(
				gomponents.Text("To"),
				html.Input(
					html.Type("datetime-local"),
					html.Name("end"),
					html.Value(endValue),
				),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text("Filter"),
			),
		),
	)
}
