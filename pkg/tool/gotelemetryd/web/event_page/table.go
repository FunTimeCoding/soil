package event_page

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(events []usage_event.Event) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(events))

	for _, e := range events {
		var shown gomponents.Node

		if e.Detail != nil {
			var parsed map[string]string

			if json.Unmarshal([]byte(*e.Detail), &parsed) == nil {
				shown = detail(parsed)
			}
		}

		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(e.CreatedAt),
				html.Td(gomponents.Text(e.Tool)),
				html.Td(gomponents.Text(e.Surface)),
				html.Td(gomponents.Text(e.Actor)),
				html.Td(gomponents.Text(e.Kind)),
				html.Td(gomponents.Text(e.Outcome)),
				html.Td(shown),
			),
		)
	}

	return table.New(
		"Time",
		"Tool",
		"Surface",
		"Actor",
		"Kind",
		"Outcome",
		"Detail",
	).
		Rows(rows).
		Empty("No events recorded.").
		Node()
}
