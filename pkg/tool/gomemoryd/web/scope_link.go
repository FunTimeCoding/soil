package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func scopeLink(
	label string,
	value string,
	current string,
	tag string,
	memoryType string,
) gomponents.Node {
	active := value == current || (value == "" && current == "all")

	if active {
		return html.Strong(gomponents.Text(label))
	}

	return html.A(
		gomponents.Attr("href", pageLink(1, tag, memoryType, value)),
		gomponents.Text(label),
	)
}
