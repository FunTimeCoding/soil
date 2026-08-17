package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func labelPips(labels []label.Label) gomponents.Node {
	var pips []gomponents.Node

	for _, l := range labels {
		pips = append(
			pips,
			html.Span(
				html.Class("label-pip"),
				html.Span(html.Class("label-key"), gomponents.Text(l.Key)),
				gomponents.Text(l.Value),
			),
		)
	}

	return html.P(gomponents.Group(pips))
}
