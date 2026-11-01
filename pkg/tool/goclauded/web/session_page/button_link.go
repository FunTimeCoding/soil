package session_page

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func buttonLink(
	path string,
	label string,
) gomponents.Node {
	return html.A(
		html.Href(path),
		html.Role("button"),
		html.Class("outline secondary"),
		gomponents.Text(label),
	)
}
