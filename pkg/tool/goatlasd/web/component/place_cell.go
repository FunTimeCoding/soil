package component

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func PlaceCell(
	kind string,
	name string,
) gomponents.Node {
	if name == "" {
		return html.Td()
	}

	return html.Td(
		html.A(html.Href(PlaceLocator(kind, name)), gomponents.Text(name)),
	)
}
