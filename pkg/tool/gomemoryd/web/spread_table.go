package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func spreadTable(summary *token_summary.Summary) gomponents.Node {
	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("")),
				html.Th(gomponents.Text("Median")),
				html.Th(gomponents.Text("Ninetieth")),
				html.Th(gomponents.Text("Ninety-ninth")),
				html.Th(gomponents.Text("Maximum")),
				html.Th(gomponents.Text("Total")),
			),
		),
		html.TBody(
			spreadRow("Block", summary.Block),
			spreadRow("Description", summary.Description),
		),
	)
}
