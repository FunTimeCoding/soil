package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func rowSpans(
	labels []string,
	values []string,
) gomponents.Node {
	var nodes []gomponents.Node

	for i, label := range labels {
		value := constant.PendingValue

		if i < len(values) {
			value = values[i]
		}

		nodes = append(
			nodes,
			html.Span(
				gomponents.Text(label),
				gomponents.Text(" "),
				html.Span(
					html.Class("board-row-value"),
					gomponents.Text(value),
				),
			),
		)
	}

	return gomponents.Group(nodes)
}
