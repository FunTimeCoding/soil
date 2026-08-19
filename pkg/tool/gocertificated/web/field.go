package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func field(
	r *http.Request,
	name string,
	label string,
	hint string,
) gomponents.Node {
	return html.Label(
		gomponents.Text(label),
		html.Input(
			html.Type("text"),
			html.Name(name),
			html.Value(r.URL.Query().Get(name)),
			html.Placeholder(hint),
		),
	)
}
