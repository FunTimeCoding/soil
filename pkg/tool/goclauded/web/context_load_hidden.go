package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func contextLoadHidden(reason error) []gomponents.Node {
	return []gomponents.Node{
		html.H4(gomponents.Text("Context loaded")),
		html.P(html.Em(gomponents.Text(reason.Error()))),
	}
}
