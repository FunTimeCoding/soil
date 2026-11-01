package memory_page

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Link(
	identifier int64,
	name string,
) gomponents.Node {
	return html.A(html.Href(Path(identifier)), gomponents.Text(name))
}
