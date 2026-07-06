package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func entriesTable(entries []entry.Entry) gomponents.Node {
	if len(entries) == 0 {
		return html.P(html.Em(gomponents.Text("No entries found.")))
	}

	var rows []gomponents.Node

	for _, e := range entries {
		target := fmt.Sprintf("detail-%d", e.ID)
		rows = append(
			rows,
			html.Tr(
				html.ID(fmt.Sprintf("row-%d", e.ID)),
				html.Class("clickable-row"),
				htmx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
				htmx.Target(fmt.Sprintf("#%s", target)),
				htmx.Swap("outerHTML"),
				html.Td(gomponents.Text(formatTime(e.Timestamp))),
				html.Td(gomponents.Text(e.Action)),
				html.Td(gomponents.Text(e.User)),
				html.Td(gomponents.Text(e.System)),
				html.Td(gomponents.Text(e.Service)),
				html.Td(gomponents.Text(truncate(e.Description, 80))),
			),
			html.Tr(
				html.ID(target),
				html.Style("display:none"),
			),
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
