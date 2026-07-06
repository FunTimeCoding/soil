package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func documentsTable(documents []store.SearchResult) gomponents.Node {
	if len(documents) == 0 {
		return html.P(gomponents.Text("No documents."))
	}

	var rows []gomponents.Node

	for _, d := range documents {
		sourceType := ""

		if d.SourceType != "" {
			sourceType = d.SourceType
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf(
								"/documents/%s",
								documentPath(d.VirtualPath),
							),
						),
						gomponents.Text(d.Title),
					),
				),
				html.Td(
					html.Small(
						gomponents.Text(d.VirtualPath),
					),
				),
				html.Td(
					html.Small(
						gomponents.Text(sourceType),
					),
				),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Title")),
				html.Th(gomponents.Text("Path")),
				html.Th(gomponents.Text("Type")),
			),
		),
		html.TBody(rows...),
	)
}
