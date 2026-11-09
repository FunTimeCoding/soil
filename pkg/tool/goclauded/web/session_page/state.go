package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/roster_card"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func State(c *roster_card.Card) gomponents.Node {
	s := c.Session
	var lines []gomponents.Node

	if s.Topic != "" {
		lines = append(
			lines,
			html.P(
				html.Class(
					join.Space(
						webConstant.CardSubtitleClass,
						webConstant.CardClampClass,
					),
				),
				gomponents.Text(s.Topic),
			),
		)
	}

	if s.Alias != nil && *s.Alias != "" {
		lines = append(
			lines,
			html.P(
				html.Class(webConstant.CardMetaClass),
				gomponents.Text(*s.Alias),
			),
		)
	}

	if s.Description != "" {
		lines = append(
			lines,
			html.P(
				html.Class(
					join.Space(
						webConstant.CardMetaClass,
						webConstant.CardClampClass,
					),
				),
				gomponents.Text(s.Description),
			),
		)
	}

	meta := []string{shortenPaths(s.Files)}
	var counts []string

	if s.TurnCount > 0 {
		counts = append(counts, fmt.Sprintf("%d turns", s.TurnCount))
	}

	if s.Lines > 0 {
		counts = append(counts, fmt.Sprintf("%d lines", s.Lines))
	}

	if s.ContextPercent > 0 {
		counts = append(counts, fmt.Sprintf("%d%% context", s.ContextPercent))
	}

	if c.Queued > 0 {
		counts = append(counts, fmt.Sprintf("%d queued", c.Queued))
	}

	counts = append(counts, time.Relative(activity(s)))
	meta = append(meta, join.Middot(counts))

	for _, line := range meta {
		if line != "" {
			lines = append(
				lines,
				html.P(
					html.Class(webConstant.CardMetaClass),
					gomponents.Text(line),
				),
			)
		}
	}

	lines = append(lines, component.LabelChips(c.Labels))

	if c.Latest != nil {
		lines = append(lines, html.P(html.Em(gomponents.Text(c.Latest.Body))))
	}

	return gomponents.Group(lines)
}
