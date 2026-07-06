package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func heatmapTable(summaries []store.Summary) gomponents.Node {
	if len(summaries) == 0 {
		return html.P(gomponents.Text("No clicks recorded yet."))
	}

	var rows []gomponents.Node

	for _, v := range summaries {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(v.Label)),
				html.Td(
					html.Class("heatmap-count"),
					gomponents.Text(fmt.Sprintf("%d", v.Count)),
				),
				html.Td(
					html.Class("heatmap-last"),
					gomponents.Text(v.Last.Format("2006-01-02 15:04")),
				),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Link")),
				html.Th(gomponents.Text("Clicks")),
				html.Th(gomponents.Text("Last clicked")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
