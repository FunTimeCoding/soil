package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func placeCell(
	kind string,
	name string,
) gomponents.Node {
	if name == "" {
		return html.Td()
	}

	return html.Td(
		html.A(html.Href(placeLocator(kind, name)), gomponents.Text(name)),
	)
}
