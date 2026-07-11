package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Alert(message string) gomponents.Node {
	return html.P(
		gomponents.Attr("role", "alert"),
		html.Class("form-alert"),
		gomponents.Text(message),
	)
}
