package web

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func timelineRow(e *timeline.Entry) gomponents.Node {
	description := timeline.FormatDescription(e)
	eventCell := []gomponents.Node{gomponents.Text(description)}

	if e.Identifier > 0 && isEditable(e.Kind) {
		eventCell = append(
			eventCell,
			gomponents.Text(" "),
			html.Span(
				html.Class("rename-icon"),
				extended.Get(fmt.Sprintf("/history/%d/edit", e.Identifier)),
				extended.Target(fmt.Sprintf("#event-%d", e.Identifier)),
				extended.Swap("innerHTML"),
				gomponents.Text("✎"),
			),
		)
	}

	return html.Tr(
		html.ID(fmt.Sprintf("event-%d", e.Identifier)),
		layout.TimeCell(library.Parse(time.RFC3339, e.Timestamp)),
		html.Td(gomponents.Group(eventCell)),
	)
}
