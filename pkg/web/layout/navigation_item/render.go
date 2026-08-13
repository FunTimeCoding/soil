package navigation_item

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (i *Item) Render(currentPath string) gomponents.Node {
	var attributes []gomponents.Node
	attributes = append(
		attributes,
		gomponents.Attr("href", i.path),
		gomponents.Text(i.label),
	)

	if i.external {
		attributes = append(attributes, gomponents.Attr("target", "_blank"))
	}

	if currentPath == i.path {
		attributes = append(attributes, gomponents.Attr("aria-current", "page"))
	}

	return html.Li(html.A(attributes...))
}
