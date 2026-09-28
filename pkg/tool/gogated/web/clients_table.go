package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func clientsTable(clients []*client.Client) gomponents.Node {
	if len(clients) == 0 {
		return html.P(gomponents.Text("No clients registered."))
	}

	var rows []gomponents.Node

	for _, c := range clients {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf("/clients/%s", c.Identifier),
						),
						gomponents.Text(truncate(c.Identifier, 8)),
					),
				),
				html.Td(gomponents.Text(c.RedirectLocators)),
				html.Td(gomponents.Text(c.GrantTypes)),
				layout.TimeCell(c.CreatedAt),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Identifier")),
				html.Th(gomponents.Text("Redirect Locators")),
				html.Th(gomponents.Text("Grant Types")),
				html.Th(gomponents.Text("Created")),
			),
		),
		html.TBody(rows...),
	)
}
