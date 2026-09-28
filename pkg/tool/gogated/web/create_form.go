package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func createForm(
	locators string,
	scopes string,
) gomponents.Node {
	if scopes == "" {
		scopes = "openid offline"
	}

	return html.Form(
		html.Method("post"),
		html.Action(constant.CreatePath),
		html.Label(
			html.For("redirect_locators"),
			gomponents.Text("Redirect Locators (one per line)"),
		),
		html.Textarea(
			html.ID("redirect_locators"),
			html.Name("redirect_locators"),
			html.Rows("3"),
			gomponents.Attr("required", ""),
			gomponents.Text(locators),
		),
		html.Label(html.For("scopes"), gomponents.Text("Scopes")),
		html.Input(
			html.Type("text"),
			html.ID("scopes"),
			html.Name("scopes"),
			html.Value(scopes),
		),
		html.Button(html.Type("submit"), gomponents.Text("Create")),
	)
}
