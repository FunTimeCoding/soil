package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func detailRow(
	label string,
	value string,
) gomponents.Node {
	return html.P(
		html.Strong(gomponents.Textf("%s: ", label)),
		gomponents.Text(value),
	)
}
