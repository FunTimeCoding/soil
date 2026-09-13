package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/web/extended"
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
		extended.StreamSwap(eventName(v.Label)),
		rowSpans(labels, values),
	)
}
