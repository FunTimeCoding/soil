package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func certificatesTable(v []record.Record) gomponents.Node {
	if len(v) == 0 {
		return html.P(html.Em(gomponents.Text("Nothing issued yet.")))
	}

	var rows []gomponents.Node

	for _, r := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(r.CommonName)),
				html.Td(gomponents.Text(r.Kind)),
				html.Td(gomponents.Text(r.Issuer)),
				layout.TimeCell(r.End),
				html.Td(gomponents.Text(state(&r))),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Common name")),
				html.Th(gomponents.Text("Kind")),
				html.Th(gomponents.Text("Issuer")),
				html.Th(gomponents.Text("Expires")),
				html.Th(gomponents.Text("State")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
