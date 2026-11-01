package dashboard_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (p *Page) Impressions() gomponents.Node {
	rows := make([]gomponents.Node, 0, len(p.view.Impressions))

	for _, i := range p.view.Impressions {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(library.Parse(time.RFC3339, i.CreatedAt)),
				html.Td(html.Em(gomponents.Text(i.Source))),
				html.Td(gomponents.Text(i.Content)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Recent Impressions")),
			table.New("Time", "Source", "Content").
				Rows(rows).
				Empty("No recent impressions.").
				Node(),
		},
	)
}
