package session_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) loads() gomponents.Node {
	heading := html.H4(gomponents.Text("Context loaded"))

	if p.view.LoadsHidden != "" {
		return gomponents.Group(
			[]gomponents.Node{
				heading,
				html.P(
					html.Class(webConstant.AbsentClass),
					gomponents.Text(p.view.LoadsHidden),
				),
			},
		)
	}

	if len(p.view.Loads) == 0 {
		return gomponents.Group(nil)
	}

	var rows []gomponents.Node
	base := memoryBase()

	for _, entry := range collapseSearches(p.view.Loads) {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(entry.OccurredAt),
				html.Td(html.Small(gomponents.Text(displayKind(&entry)))),
				referenceCell(&entry, base),
				html.Td(gomponents.Text(entry.Name)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			heading,
			table.New("Time", "Kind", "Reference", "Name").Rows(rows).Node(),
		},
	)
}
