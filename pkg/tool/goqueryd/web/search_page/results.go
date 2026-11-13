package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	"github.com/funtimecoding/soil/pkg/web/component/highlight"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Results(
	results []result.Search,
	all []result.Facet,
) gomponents.Node {
	terms := p.terms()
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
				html.Td(html.Small(highlight.New(m.Snippet, terms...).Node())),
				html.Td(gomponents.Textf("%.2f", m.Score)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			p.filters(),
			count(len(results), p.collection),
			p.facets(all),
			table.New("Title", "Path", "Snippet", "Score").
				Numeric(3).
				Rows(rows).
				Empty("No results.").
				Node(),
		},
	)
}
