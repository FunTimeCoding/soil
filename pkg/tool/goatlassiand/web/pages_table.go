package web

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func pagesTable(pages []*page.Page) gomponents.Node {
	if len(pages) == 0 {
		return html.P(gomponents.Text("None."))
	}

	rows := make([]gomponents.Node, 0, len(pages))

	for _, p := range pages {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(p.Link),
						html.Target("_blank"),
						gomponents.Text(p.Name),
					),
				),
				layout.TimeCell(p.Raw.Version.CreatedAt),
			),
		)
	}

	return html.Table(html.Class("pages-table"), html.TBody(rows...))
}
