package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func raidsTable(rows []store.RaidRow) gomponents.Node {
	if len(rows) == 0 {
		return html.P(html.Em(gomponents.Text("No raids created yet.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Date")),
				html.Th(gomponents.Text("Fights")),
				html.Th(gomponents.Text(constant.PlayersTitle)),
			),
		),
		html.TBody(
			gomponents.Map(
				rows,
				func(r store.RaidRow) gomponents.Node {
					return html.Tr(
						html.Td(
							html.A(
								html.Href(
									fmt.Sprintf("/raids/%d", r.ID),
								),
								gomponents.Text(r.Name),
							),
						),
						html.Td(
							html.Class(layout.TimeCellClass),
							gomponents.Text(r.Date.Format(time.DateYear)),
						),
						html.Td(gomponents.Textf("%d", r.Fights)),
						html.Td(gomponents.Textf("%d", r.Players)),
					)
				},
			),
		),
	)
}
