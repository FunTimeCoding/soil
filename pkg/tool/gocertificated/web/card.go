package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func card(
	label string,
	value int,
) gomponents.Node {
	return html.Article(
		html.Header(gomponents.Text(label)),
		html.P(gomponents.Textf("%d", value)),
	)
}
