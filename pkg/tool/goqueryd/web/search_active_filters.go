package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"sort"
)

func searchActiveFilters(
	metadata map[string]string,
	query string,
	collection string,
) gomponents.Node {
	if len(metadata) == 0 {
		return gomponents.Group(nil)
	}

	keys := make([]string, 0, len(metadata))

	for k := range metadata {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var chips []gomponents.Node

	for _, k := range keys {
		without := copyWithout(metadata, k)
		chips = append(
			chips,
			html.Span(
				gomponents.Textf("%s=%s ", k, metadata[k]),
				html.A(
					gomponents.Attr(
						"hx-get",
						searchLink(query, collection, without),
					),
					gomponents.Attr("hx-target", "#search-results"),
					gomponents.Attr("hx-swap", "innerHTML"),
					gomponents.Text("×"),
				),
				gomponents.Text("  "),
			),
		)
	}

	return html.P(
		html.Small(gomponents.Text("Filters: "), gomponents.Group(chips)),
	)
}
