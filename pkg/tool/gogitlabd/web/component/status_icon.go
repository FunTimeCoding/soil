package component

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func StatusIcon(status string) gomponents.Node {
	return html.Img(
		html.Class("status-icon"),
		html.Src(statusSource(status)),
		html.Alt(status),
		html.Title(status),
	)
}
