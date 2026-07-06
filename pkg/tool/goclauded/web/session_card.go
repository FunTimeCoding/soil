package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func sessionCard(
	s *session.Session,
	lines int,
	labels []label.Label,
	latestPulse *pulse.Pulse,
) gomponents.Node {
	var details []gomponents.Node

	if s.Alias != nil {
		details = append(
			details,
			html.P(html.Strong(gomponents.Text(*s.Alias))),
		)
	}

	if s.Topic != "" {
		details = append(
			details,
			html.P(gomponents.Text(s.Topic)),
		)
	}

	if s.Description != "" {
		details = append(
			details,
			html.P(
				html.Small(gomponents.Text(s.Description)),
			),
		)
	}

	if s.Files != "" {
		details = append(
			details,
			html.P(
				html.Small(gomponents.Text(shortenPaths(s.Files))),
			),
		)
	}

	var metadata []string

	if s.TurnCount > 0 {
		metadata = append(metadata, fmt.Sprintf("%d turns", s.TurnCount))
	}

	if lines > 0 {
		metadata = append(metadata, fmt.Sprintf("%d lines", lines))
	}

	metadata = append(metadata, relativeTime(s.LastSeen))
	details = append(
		details,
		html.P(
			html.Small(gomponents.Text(strings.Join(metadata, " · "))),
		),
	)

	if len(labels) > 0 {
		var pips []gomponents.Node

		for _, l := range labels {
			pips = append(
				pips,
				html.Span(
					html.Class("label-pip"),
					gomponents.Text(
						fmt.Sprintf("(%s:%s)", l.Key, l.Value),
					),
				),
			)
		}

		details = append(details, html.P(gomponents.Group(pips)))
	}

	if latestPulse != nil {
		details = append(
			details,
			html.P(
				html.Em(
					html.Class("pulse-line"),
					gomponents.Text(latestPulse.Body),
				),
			),
		)
	}

	return html.A(
		gomponents.Attr(
			"href",
			fmt.Sprintf("/sessions/%s", s.Identifier),
		),
		html.Class("session-card"),
		html.H4(
			statusDot(s.LastSeen),
			gomponents.Text(s.Name),
		),
		gomponents.Group(details),
	)
}
