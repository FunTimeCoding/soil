package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Results(
	results []result.Search,
	all []result.Facet,
	query string,
	collection string,
	metadata map[string]string,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(results))

	for _, m := range results {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(document_page.Path(m.VirtualPath)),
						gomponents.Text(m.Title),
					),
				),
				html.Td(html.Small(gomponents.Text(m.VirtualPath))),
				html.Td(html.Small(gomponents.Text(m.Snippet))),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%.2f", m.Score),
				),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			filters(query, collection, metadata),
			count(len(results), collection),
			facets(all, query, collection, metadata),
			table.New("Title", "Path", "Snippet", "Score").
				Numeric(3).
				Rows(rows).
				Empty("No results.").
				Node(),
		},
	)
}
