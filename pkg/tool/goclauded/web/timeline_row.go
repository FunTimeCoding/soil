package web

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func timelineRow(e *timeline.Entry) gomponents.Node {
	t, f := time.Parse(time.RFC3339, e.Timestamp)
	display := e.Timestamp

	if f == nil {
		display = library.FormatCompact(t)
	}

	description := timeline.FormatDescription(e)
	eventCell := []gomponents.Node{gomponents.Text(description)}

	if e.Identifier > 0 && isEditable(e.Kind) {
		eventCell = append(
			eventCell,
			gomponents.Text(" "),
			html.Span(
				html.Class("rename-icon"),
				gomponents.Attr(
					"hx-get",
					fmt.Sprintf("/history/%d/edit", e.Identifier),
				),
				gomponents.Attr(
					"hx-target",
					fmt.Sprintf("#event-%d", e.Identifier),
				),
				gomponents.Attr("hx-swap", "innerHTML"),
				gomponents.Text("✎"),
			),
		)
	}

	return html.Tr(
		html.ID(fmt.Sprintf("event-%d", e.Identifier)),
		html.Td(gomponents.Text(display)),
		html.Td(gomponents.Group(eventCell)),
	)
}
