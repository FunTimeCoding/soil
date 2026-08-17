package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func contextLoadSection(loads []context_load.Load) []gomponents.Node {
	if len(loads) == 0 {
		return nil
	}

	var rows []gomponents.Node

	for _, entry := range collapseSearches(loads) {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(entry.OccurredAt),
				html.Td(html.Small(gomponents.Text(displayKind(&entry)))),
				html.Td(html.Small(gomponents.Text(entry.Reference))),
				html.Td(gomponents.Text(entry.Name)),
			),
		)
	}

	return []gomponents.Node{
		html.H4(gomponents.Text("Context loaded")),
		html.Table(
			html.THead(
				html.Tr(
					html.Th(gomponents.Text("Time")),
					html.Th(gomponents.Text("Kind")),
					html.Th(gomponents.Text("Reference")),
					html.Th(gomponents.Text("Name")),
				),
			),
			html.TBody(rows...),
		),
	}
}
