package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func cruiseOption(
	value string,
	text string,
	current string,
) gomponents.Node {
	return html.Option(
		html.Value(value),
		gomponents.If(value == current, html.Selected()),
		gomponents.Text(text),
	)
}
