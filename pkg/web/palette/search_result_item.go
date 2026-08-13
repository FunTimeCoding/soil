package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func searchResultItem(item SearchItem) gomponents.Node {
	var children []gomponents.Node
	children = append(
		children,
		html.Span(html.Class("palette-label"), gomponents.Text(item.Label)),
	)

	if item.Description != "" {
		children = append(
			children,
			html.Span(
				html.Class("palette-description"),
				gomponents.Text(item.Description),
			),
		)
	}

	if item.Category != "" {
		children = append(
			children,
			html.Span(
				html.Class("palette-category"),
				gomponents.Text(item.Category),
			),
		)
	}

	return html.A(
		html.Class("palette-item"),
		gomponents.Attr("href", item.Path),
		gomponents.Attr("data-palette-item", ""),
		gomponents.Group(children),
	)
}
