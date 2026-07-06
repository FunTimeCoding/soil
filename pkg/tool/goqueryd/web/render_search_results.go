package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) renderSearchResults(
	results []store.SearchResult,
	facets []store.Facet,
	degraded bool,
	query string,
	collection string,
	metadata map[string]string,
) []gomponents.Node {
	activeFilters := searchActiveFilters(metadata, query, collection)
	resultCount := searchResultCount(len(results), collection)
	var content []gomponents.Node
	content = append(content, activeFilters, resultCount)

	if degraded {
		content = append(
			content,
			html.P(
				html.Small(
					html.Em(
						gomponents.Text("degraded to keyword search"),
					),
				),
			),
		)
	}

	content = append(
		content,
		searchFacets(facets, query, collection, metadata),
	)

	if len(results) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No results.")),
		)
	} else {
		var rows []gomponents.Node

		for _, m := range results {
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.A(
							gomponents.Attr(
								"href",
								fmt.Sprintf(
									"/documents/%s",
									documentPath(m.VirtualPath),
								),
							),
							gomponents.Text(m.Title),
						),
					),
					html.Td(
						html.Small(
							gomponents.Text(m.VirtualPath),
						),
					),
					html.Td(
						html.Small(
							gomponents.Text(m.Snippet),
						),
					),
					html.Td(
						gomponents.Textf("%.2f", m.Score),
					),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Title")),
						html.Th(gomponents.Text("Path")),
						html.Th(gomponents.Text("Snippet")),
						html.Th(gomponents.Text("Score")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	return content
}
