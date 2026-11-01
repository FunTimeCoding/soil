package client_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(clients []*client.Client) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(clients))

	for _, c := range clients {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(Path(c.Identifier)),
						gomponents.Text(component.Truncate(c.Identifier, 8)),
					),
				),
				html.Td(gomponents.Text(c.RedirectLocators)),
				html.Td(gomponents.Text(c.GrantTypes)),
				layout.TimeCell(c.CreatedAt),
			),
		)
	}

	return table.New(
		"Identifier",
		"Redirect Locators",
		"Grant Types",
		"Created",
	).
		Rows(rows).
		Empty("No clients registered.").
		Node()
}
