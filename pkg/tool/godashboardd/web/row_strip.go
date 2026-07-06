package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func rowStrip(
	v *board.Entry,
	values []string,
) gomponents.Node {
	labels := rowLabels(v)

	if len(labels) == 0 {
		return nil
	}

	return html.Div(
		html.Class("board-rows"),
		gomponents.Attr("sse-swap", eventName(v.Label)),
		rowSpans(labels, values),
	)
}
