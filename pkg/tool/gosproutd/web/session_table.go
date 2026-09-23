package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) sessionTable() gomponents.Node {
	var rows []gomponents.Node

	for _, v := range s.service.Sessions() {
		bumped := gomponents.Text("")

		if v.Bumped > 0 {
			bumped = html.Strong(gomponents.Text(fmt.Sprintf("%d", v.Bumped)))
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf(
								"%s?%s=%s",
								constant.SessionPath,
								constant.SessionParameter,
								v.Name,
							),
						),
						gomponents.Text(v.Name),
					),
				),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Open))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Pending))),
				html.Td(gomponents.Text(fmt.Sprintf("%d", v.Cleared))),
				html.Td(bumped),
				layout.TimeCell(v.LastAt),
			),
		)
	}

	if rows == nil {
		return html.P(
			html.Small(gomponents.Text("No session has pushed a decision.")),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Session")),
				html.Th(gomponents.Text("Open")),
				html.Th(gomponents.Text("Waiting on me")),
				html.Th(gomponents.Text("Cleared")),
				html.Th(gomponents.Text("Bumped")),
				html.Th(gomponents.Text("Last")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
