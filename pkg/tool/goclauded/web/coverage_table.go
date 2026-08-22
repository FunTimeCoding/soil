package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageTable(servers []*coverage.Server) gomponents.Node {
	if len(servers) == 0 {
		return html.P(gomponents.Text("No coverage data."))
	}

	var rows []gomponents.Node

	for _, e := range servers {
		rows = append(rows, coverageRow(e))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Server")),
				html.Th(gomponents.Text("Coverage 30d")),
				html.Th(gomponents.Text("Coverage all")),
				html.Th(gomponents.Text("Calls 30d")),
				html.Th(gomponents.Text("Calls all")),
				html.Th(gomponents.Text("Last used")),
			),
		),
		html.TBody(rows...),
	)
}
