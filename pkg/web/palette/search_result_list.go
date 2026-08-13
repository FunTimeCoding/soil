package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SearchResultList(items []SearchItem) gomponents.Node {
	if len(items) == 0 {
		return html.Div(
			html.ID("palette-search-results"),
			html.Div(html.Class("palette-empty"), gomponents.Text("No matches")),
		)
	}

	var nodes []gomponents.Node

	for _, item := range items {
		nodes = append(nodes, searchResultItem(item))
	}

	return html.Div(html.ID("palette-search-results"), gomponents.Group(nodes))
}
