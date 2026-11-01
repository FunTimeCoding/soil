package client_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	"github.com/funtimecoding/soil/pkg/web/component/panel"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Created(
	identifier string,
	secret string,
) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H1(gomponents.Text("Client Created")),
			panel.New("Credentials").
				Content(
					fact.New().
						RowNode(
							"Client Identifier",
							html.Code(gomponents.Text(identifier)),
						).
						RowNode(
							"Client Secret",
							html.Code(gomponents.Text(secret)),
						).
						Node(),
				).
				Footer(
					html.P(
						html.Class(constant.WarningClass),
						gomponents.Text(
							"Save the secret now. It will not be shown again.",
						),
					),
				).
				Node(),
			load.New(constant.ClientsPath, "Back to Clients").
				Button().
				Class("secondary").
				Node(),
		},
	)
}
