package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sessionsTable(
	sessions []*authentication_session.AuthenticationSession,
	current string,
) gomponents.Node {
	if len(sessions) == 0 {
		return html.P(gomponents.Text("No active sessions."))
	}

	var rows []gomponents.Node

	for _, session := range sessions {
		label := truncate(session.Identifier, 8)

		if session.Identifier == current {
			label = fmt.Sprintf("%s (this browser)", label)
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(label)),
				html.Td(gomponents.Text(session.UserAgent)),
				html.Td(gomponents.Text(session.Address)),
				layout.TimeCell(session.AuthenticatedAt),
				layout.TimeCell(session.LastUsedAt),
				html.Td(
					html.Form(
						html.Method("post"),
						html.Action(
							fmt.Sprintf(
								"/sessions/%s/delete",
								session.Identifier,
							),
						),
						html.Button(
							html.Type("submit"),
							gomponents.Attr("class", "secondary"),
							gomponents.Text("Revoke"),
						),
					),
				),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Identifier")),
				html.Th(gomponents.Text("User Agent")),
				html.Th(gomponents.Text("Address")),
				html.Th(gomponents.Text("Signed In")),
				html.Th(gomponents.Text("Last Used")),
				html.Th(gomponents.Text("")),
			),
		),
		html.TBody(rows...),
	)
}
