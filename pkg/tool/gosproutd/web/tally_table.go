package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) tallyTable() gomponents.Node {
	var rows []gomponents.Node

	for _, v := range s.service.Tallies() {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(v.Session)),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Emitted))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Answered))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Bypassed))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Defaulted))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Declined))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Irrelevant))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Postponed))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Abandoned))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Bumped))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Turns))),
			),
		)
	}

	if rows == nil {
		return html.P(html.Small(gomponents.Text("Nothing counted yet.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Session")),
				html.Th(gomponents.Text("Emitted")),
				html.Th(gomponents.Text("Answered")),
				html.Th(gomponents.Text("In chat")),
				html.Th(gomponents.Text("Defaulted")),
				html.Th(gomponents.Text("Declined")),
				html.Th(gomponents.Text("Irrelevant")),
				html.Th(gomponents.Text("Postponed")),
				html.Th(gomponents.Text("Abandoned")),
				html.Th(gomponents.Text("Bumped")),
				html.Th(gomponents.Text("Turns")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
