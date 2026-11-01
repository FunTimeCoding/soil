package session_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) completions() gomponents.Node {
	completions := p.view.Detail.Completions

	if len(completions) == 0 {
		return gomponents.Group(nil)
	}

	rows := make([]gomponents.Node, 0, len(completions))

	for _, c := range completions {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(c.CreatedAt),
				html.Td(badge.New(c.Kind).Quiet().Node()),
				html.Td(gomponents.Text(c.Topic)),
				html.Td(html.Small(gomponents.Text(c.Summary))),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H4(gomponents.Text("Completions")),
			table.New("Time", "Kind", "Topic", "Summary").Rows(rows).Node(),
		},
	)
}
