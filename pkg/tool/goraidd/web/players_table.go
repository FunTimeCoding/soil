package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func playersTable(rows []store.AttendanceRow) gomponents.Node {
	if len(rows) == 0 {
		return html.P(html.Em(gomponents.Text("No player data yet.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Account")),
				html.Th(gomponents.Text("Characters")),
				html.Th(gomponents.Text("Fights")),
				html.Th(gomponents.Text("Attendance")),
			),
		),
		html.TBody(
			gomponents.Map(
				rows,
				func(r store.AttendanceRow) gomponents.Node {
					pct := 0

					if r.Available > 0 {
						pct = r.Fights * 100 / r.Available
					}

					return html.Tr(
						html.Td(
							html.A(
								html.Href(
									fmt.Sprintf(
										"/players/%s",
										r.Account,
									),
								),
								gomponents.Text(r.Account),
							),
						),
						html.Td(gomponents.Text(r.Characters)),
						html.Td(
							gomponents.Textf(
								"%d / %d",
								r.Fights,
								r.Available,
							),
						),
						html.Td(gomponents.Textf("%d%%", pct)),
					)
				},
			),
		),
	)
}
