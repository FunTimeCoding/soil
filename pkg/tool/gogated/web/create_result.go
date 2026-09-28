package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func createResult(
	clientIdentifier string,
	clientSecret string,
) gomponents.Node {
	return html.Div(
		html.Class("credential-card"),
		html.P(html.Strong(gomponents.Text("Client Identifier"))),
		html.P(html.Code(gomponents.Text(clientIdentifier))),
		html.P(html.Strong(gomponents.Text("Client Secret"))),
		html.P(html.Code(gomponents.Text(clientSecret))),
		html.P(
			html.Class(constant.WarningClass),
			gomponents.Text("Save the secret now. It will not be shown again."),
		),
		html.A(
			gomponents.Attr("href", constant.ClientsPath),
			gomponents.Attr("role", "button"),
			html.Class("secondary"),
			gomponents.Text("Back to Clients"),
		),
	)
}
