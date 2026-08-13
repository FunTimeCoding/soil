package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func resultList(results []Result) gomponents.Node {
	if len(results) == 0 {
		return html.Div(
			html.Class("palette-empty"),
			gomponents.Text("No matches"),
		)
	}

	var items []gomponents.Node

	for _, r := range results {
		items = append(items, resultItem(r))
	}

	return html.Div(html.ID("palette-results"), gomponents.Group(items))
}
