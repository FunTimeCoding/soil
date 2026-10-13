package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func requestTable(
	requests []*request_entry.Entry,
	multiple bool,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(requests))

	for _, entry := range requests {
		r := entry.Request
		var created time.Time

		if r.Create != nil {
			created = *r.Create
		}

		rows = append(
			rows,
			html.Tr(
				gomponents.If(
					multiple,
					html.Td(
						html.Class("reference"),
						gomponents.Text(entry.Instance),
					),
				),
				html.Td(
					html.A(
						html.Href(r.Link),
						html.Target("_blank"),
						gomponents.Textf("!%d", r.Identifier),
					),
				),
				html.Td(gomponents.Text(r.Title)),
				html.Td(html.Class("reference"), gomponents.Text(r.State)),
				layout.TimeCell(created),
			),
		)
	}

	return html.Table(
		html.Class("board-table"),
		html.THead(
			html.Tr(
				gomponents.If(multiple, html.Th(gomponents.Text("Instance"))),
				html.Th(gomponents.Text("Request")),
				html.Th(gomponents.Text("Title")),
				html.Th(gomponents.Text("State")),
				html.Th(gomponents.Text("Created")),
			),
		),
		html.TBody(rows...),
	)
}
