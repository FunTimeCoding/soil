package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"sort"
)

func searchFacets(
	facets []store.Facet,
	query string,
	collection string,
	activeMetadata map[string]string,
) gomponents.Node {
	if len(facets) == 0 {
		return gomponents.Group(nil)
	}

	var sections []gomponents.Node

	for _, f := range facets {
		if len(f.Values) == 0 {
			sections = append(
				sections,
				html.P(
					html.Small(
						gomponents.Textf(
							"%s: %d distinct",
							f.Key,
							f.Distinct,
						),
					),
				),
			)

			continue
		}

		keys := make([]string, 0, len(f.Values))

		for k := range f.Values {
			keys = append(keys, k)
		}

		sort.Strings(keys)
		var items []gomponents.Node

		for _, value := range keys {
			count := f.Values[value]

			if activeMetadata[f.Key] == value {
				without := copyWithout(activeMetadata, f.Key)
				items = append(
					items,
					html.Span(
						html.Strong(
							gomponents.Textf("%s (%d)", value, count),
						),
						gomponents.Text(" "),
						html.A(
							gomponents.Attr(
								"hx-get",
								searchLink(query, collection, without),
							),
							gomponents.Attr("hx-target", "#search-results"),
							gomponents.Attr("hx-swap", "innerHTML"),
							gomponents.Text("×"),
						),
					),
				)
			} else {
				with := copyWith(activeMetadata, f.Key, value)
				items = append(
					items,
					html.A(
						gomponents.Attr(
							"hx-get",
							searchLink(query, collection, with),
						),
						gomponents.Attr("hx-target", "#search-results"),
						gomponents.Attr("hx-swap", "innerHTML"),
						gomponents.Textf("%s (%d)", value, count),
					),
				)
			}

			items = append(items, gomponents.Text(" "))
		}

		sections = append(
			sections,
			html.P(
				html.Small(
					html.Strong(gomponents.Textf("%s: ", f.Key)),
					gomponents.Group(items),
				),
			),
		)
	}

	return gomponents.Group(sections)
}
