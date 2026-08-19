package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func authoritiesTable(v []record.Record) gomponents.Node {
	if len(v) == 0 {
		return html.P(
			html.Em(
				gomponents.Text(
					"No authority yet. Create the root to begin the chain.",
				),
			),
		)
	}

	var rows []gomponents.Node

	for _, r := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(r.Name)),
				html.Td(gomponents.Text(r.Kind)),
				html.Td(gomponents.Text(r.CommonName)),
				html.Td(gomponents.Text(permitted(&r))),
				layout.TimeCell(r.End),
				html.Td(gomponents.Text(published(&r))),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Kind")),
				html.Th(gomponents.Text("Common name")),
				html.Th(gomponents.Text("Permits")),
				html.Th(gomponents.Text("Expires")),
				html.Th(gomponents.Text("Published")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
