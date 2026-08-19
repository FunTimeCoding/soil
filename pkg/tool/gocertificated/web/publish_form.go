package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func publishForm(v []*publish.Change) gomponents.Node {
	if len(v) == 0 {
		return html.P(html.Em(gomponents.Text("Everything is published.")))
	}

	var item []gomponents.Node

	for _, c := range v {
		item = append(item, html.Li(gomponents.Text(c.Path)))
	}

	return html.Div(
		html.H2(gomponents.Text(constant.PublishTitle)),
		html.Ul(gomponents.Group(item)),
		html.Form(
			html.Method(http.MethodPost),
			html.Action(constant.PublishPath),
			html.Button(
				html.Type("submit"),
				gomponents.Textf("Commit %d files", len(v)),
			),
		),
	)
}
