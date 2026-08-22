package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageToolTable(e *coverage.Server) gomponents.Node {
	var rows []gomponents.Node

	for _, t := range e.Tools {
		rows = append(rows, coverageToolRow(e, t))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Tool")),
				html.Th(gomponents.Text("State")),
				html.Th(gomponents.Text("Calls 30d")),
				html.Th(gomponents.Text("Calls all")),
				html.Th(gomponents.Text("Last used")),
			),
		),
		html.TBody(rows...),
	)
}
