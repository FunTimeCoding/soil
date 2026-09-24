package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func requestTable(requests []*merge_request.Request) gomponents.Node {
	if len(requests) == 0 {
		return html.P(gomponents.Text(constant.RequestEmpty))
	}

	rows := make([]gomponents.Node, 0, len(requests))

	for _, r := range requests {
		var created time.Time

		if r.Create != nil {
			created = *r.Create
		}

		rows = append(
			rows,
			html.Tr(
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
				html.Th(gomponents.Text("Request")),
				html.Th(gomponents.Text("Title")),
				html.Th(gomponents.Text("State")),
				html.Th(gomponents.Text("Created")),
			),
		),
		html.TBody(rows...),
	)
}
