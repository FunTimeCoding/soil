package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func searchForm(name string) gomponents.Node {
	return html.Form(
		html.Class(constant.SearchFormClass),
		html.Method(http.MethodGet),
		html.Action(constant.PlacementsPath),
		html.Input(
			html.Type("search"),
			html.Name(constant.SearchParameter),
			html.Value(name),
			html.Placeholder(constant.SearchLabel),
			html.AutoFocus(),
		),
		html.Button(
			html.Type("submit"),
			gomponents.Text(constant.SearchAction),
		),
	)
}
