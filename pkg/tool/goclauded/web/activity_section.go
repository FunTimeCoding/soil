package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) activitySection(kinds []string) gomponents.Node {
	a := argument.Timeline{Limit: 10}

	if len(kinds) > 0 {
		a.Kinds = kinds
	}

	entries, e := s.service.Timeline(a)
	errors.PanicOnError(e)

	if len(entries) == 0 {
		return html.P(gomponents.Text("No recent activity."))
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
