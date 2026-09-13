package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func entry(s *enriched_session.Session) gomponents.Node {
	name := s.Slug

	if s.Alias != "" {
		name = s.Alias
	}

	if name == "" {
		name = "unnamed"
	}

	var nodes []gomponents.Node
	nodes = append(
		nodes,
		html.Div(
			html.Class("entry-name"),
			extended.Get(fmt.Sprintf("/conversations/%s", s.Identifier)),
			extended.Target("#panel"),
			html.Span(gomponents.Text(name)),
			html.Span(
				html.Class("rename-icon"),
				extended.Get(
					fmt.Sprintf("/conversations/%s/edit", s.Identifier),
				),
				extended.Target("#panel"),
				gomponents.Attr("onclick", "event.stopPropagation()"),
				gomponents.Text("✎"),
			),
		),
	)

	if s.Description != "" {
		nodes = append(
			nodes,
			html.Small(
				html.Class("entry-description"),
				gomponents.Text(s.Description),
			),
		)
	}

	nodes = append(
		nodes,
		html.Small(gomponents.Text(relativeTimestamp(s.Timestamp))),
	)

	return html.Div(
		html.ID(fmt.Sprintf("entry-%s", s.Identifier)),
		html.Class("sidebar-entry"),
		gomponents.Group(nodes),
	)
}
