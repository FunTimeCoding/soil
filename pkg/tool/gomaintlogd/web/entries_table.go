package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func entriesTable(entries []entry.Entry) gomponents.Node {
	if len(entries) == 0 {
		return html.P(html.Em(gomponents.Text("No entries found.")))
	}

	var rows []gomponents.Node

	for _, e := range entries {
		target := fmt.Sprintf("detail-%d", e.Identifier)
		rows = append(
			rows,
			html.Tr(
				html.ID(fmt.Sprintf("row-%d", e.Identifier)),
				html.Class("clickable-row"),
				extended.Get(
					fragmentLocator(constant.DetailPath, e.Identifier),
				),
				extended.Target(fmt.Sprintf("#%s", target)),
				extended.Swap(webConstant.SwapOuter),
				layout.TimeCell(e.Timestamp),
				html.Td(gomponents.Text(e.Action)),
				html.Td(gomponents.Text(e.User)),
				html.Td(gomponents.Text(e.System)),
				html.Td(gomponents.Text(e.Service)),
				html.Td(gomponents.Text(truncate(e.Description, 80))),
			),
			html.Tr(html.ID(target), html.Style("display:none")),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Timestamp")),
				html.Th(gomponents.Text("Action")),
				html.Th(gomponents.Text("User")),
				html.Th(gomponents.Text("System")),
				html.Th(gomponents.Text("Service")),
				html.Th(gomponents.Text("Description")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
