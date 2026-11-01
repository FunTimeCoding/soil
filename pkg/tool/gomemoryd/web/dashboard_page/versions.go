package dashboard_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (p *Page) Versions() gomponents.Node {
	rows := make([]gomponents.Node, 0, len(p.view.Versions))

	for _, v := range p.view.Versions {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(library.Parse(time.RFC3339, v.ChangedAt)),
				html.Td(gomponents.Text(v.ChangeType)),
				html.Td(
					html.A(
						html.Href(memory_page.Path(v.MemoryIdentifier)),
						gomponents.Text(v.Name),
					),
				),
				html.Td(html.Small(gomponents.Text(v.Source))),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Recent Changes")),
			table.New("Time", "Action", "Memory", "Source").
				Rows(rows).
				Empty("No recent changes.").
				Node(),
		},
	)
}
