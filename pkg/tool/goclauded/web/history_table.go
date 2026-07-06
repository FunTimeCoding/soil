package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/timeline"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) historyTable(entries []*timeline.Entry) gomponents.Node {
	if len(entries) == 0 {
		return html.P(gomponents.Text("No events recorded."))
	}

	var rows []gomponents.Node

	for _, e := range entries {
		rows = append(rows, timelineRow(e))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Time")),
				html.Th(gomponents.Text("Event")),
			),
		),
		html.TBody(rows...),
	)
}
