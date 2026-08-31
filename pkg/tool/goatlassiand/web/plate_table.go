package web

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func plateTable(issues []*issue.Issue) gomponents.Node {
	if len(issues) == 0 {
		return html.P(gomponents.Text("Nothing on the plate."))
	}

	rows := make([]gomponents.Node, 0, len(issues))

	for _, i := range issues {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Span(
						html.Class(join.Space("status-dot", statusClass(i))),
						html.Title(i.Status),
					),
				),
				html.Td(
					html.A(
						html.Href(i.Link),
						html.Target("_blank"),
						gomponents.Text(i.Key),
					),
				),
				html.Td(
					html.Class(constant.PageStatus),
					gomponents.Text(i.Status),
				),
				html.Td(gomponents.Text(i.Summary)),
				layout.TimeCell(i.ChangeTime()),
			),
		)
	}

	return html.Table(
		html.Class("plate-table"),
		html.THead(
			html.Tr(
				html.Th(),
				html.Th(gomponents.Text("Key")),
				html.Th(gomponents.Text("Status")),
				html.Th(gomponents.Text("Summary")),
				html.Th(gomponents.Text("Changed")),
			),
		),
		html.TBody(rows...),
	)
}
