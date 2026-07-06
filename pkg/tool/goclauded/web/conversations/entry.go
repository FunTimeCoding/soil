package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
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
			gomponents.Attr(
				"hx-get",
				fmt.Sprintf("/conversations/%s", s.Identifier),
			),
			gomponents.Attr("hx-target", "#panel"),
			html.Span(gomponents.Text(name)),
			html.Span(
				html.Class("rename-icon"),
				gomponents.Attr(
					"hx-get",
					fmt.Sprintf("/conversations/%s/edit", s.Identifier),
				),
				gomponents.Attr("hx-target", "#panel"),
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
