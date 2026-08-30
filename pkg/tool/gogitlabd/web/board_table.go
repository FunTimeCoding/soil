package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func (*Server) boardTable(entries []*board_entry.Entry) gomponents.Node {
	if len(entries) == 0 {
		return html.P(
			gomponents.Text("No pipelines seen yet - first poll pending."),
		)
	}

	namespace := sharedNamespace(entries)
	rows := make([]gomponents.Node, 0, len(entries))

	for _, entry := range entries {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Img(
						html.Class("status-icon"),
						html.Src(statusIcon(entry.Status)),
						html.Alt(entry.Status),
						html.Title(entry.Status),
					),
				),
				html.Td(
					html.A(
						html.Href(entry.ProjectLink),
						html.Target("_blank"),
						gomponents.Text(
							strings.TrimPrefix(entry.Project, namespace),
						),
					),
				),
				html.Td(
					html.Class("reference"),
					gomponents.Text(
						strings.TrimPrefix(
							entry.Reference,
							constant.RenovatePrefix,
						),
					),
				),
				html.Td(
					html.A(
						html.Href(entry.Link),
						html.Target("_blank"),
						gomponents.Textf("#%d", entry.Identifier),
					),
				),
				layout.TimeCell(entry.Updated),
			),
		)
	}

	return html.Table(
		html.Class("board-table"),
		html.THead(
			html.Tr(
				html.Th(),
				html.Th(gomponents.Text("Project")),
				html.Th(gomponents.Text("Branch")),
				html.Th(gomponents.Text("Pipeline")),
				html.Th(gomponents.Text("Updated")),
			),
		),
		html.TBody(rows...),
	)
}
