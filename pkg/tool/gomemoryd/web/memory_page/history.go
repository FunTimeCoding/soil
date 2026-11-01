package memory_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (p *Page) history() gomponents.Node {
	rows := make([]gomponents.Node, 0, len(p.view.History))

	for _, v := range p.view.History {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(library.Parse(time.RFC3339, v.ChangedAt)),
				html.Td(gomponents.Text(v.ChangeType)),
				html.Td(html.Small(gomponents.Text(v.Source))),
				html.Td(html.Small(gomponents.Text(v.Description))),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H4(gomponents.Text("History")),
			table.New("Time", "Action", "Source", "Description").
				Rows(rows).
				Empty("No history.").
				Node(),
		},
	)
}
