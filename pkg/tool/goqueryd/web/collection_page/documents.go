package collection_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Documents(documents []result.Search) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(documents))

	for _, d := range documents {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(document_page.Path(d.VirtualPath)),
						gomponents.Text(d.Title),
					),
				),
				html.Td(html.Small(gomponents.Text(d.VirtualPath))),
				html.Td(html.Small(gomponents.Text(d.SourceType))),
			),
		)
	}

	return table.New("Title", "Path", "Type").
		Rows(rows).
		Empty("No documents.").
		Node()
}
