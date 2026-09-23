package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) decisionTable(session string) gomponents.Node {
	var rows []gomponents.Node

	for _, v := range s.service.Decisions(session) {
		rows = append(rows, decisionRow(v))
	}

	if rows == nil {
		return html.P(
			html.Small(gomponents.Text("This session has pushed nothing.")),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Decision")),
				html.Th(gomponents.Text("Frames")),
				html.Th(gomponents.Text("State")),
				html.Th(gomponents.Text("Pushed")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
