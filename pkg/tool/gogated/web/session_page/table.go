package session_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(
	sessions []*authentication_session.AuthenticationSession,
	current string,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(sessions))

	for _, s := range sessions {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					gomponents.Text(component.Truncate(s.Identifier, 8)),
					gomponents.If(
						s.Identifier == current,
						badge.New("this browser").Quiet().Node(),
					),
				),
				html.Td(gomponents.Text(s.UserAgent)),
				html.Td(gomponents.Text(s.Address)),
				layout.TimeCell(s.AuthenticatedAt),
				layout.TimeCell(s.LastUsedAt),
				html.Td(
					action.New(
						join.Empty(
							constant.SessionsPath,
							stringsConstant.Slash,
							s.Identifier,
							constant.DeleteSuffix,
						),
						"Revoke",
					).
						ButtonClass("secondary").
						Confirm("Revoke this session?").
						Small().
						Node(),
				),
			),
		)
	}

	return table.New(
		"Identifier",
		"User Agent",
		"Address",
		"Signed In",
		"Last Used",
		"",
	).
		Rows(rows).
		Empty("No active sessions.").
		Node()
}
