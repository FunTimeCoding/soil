package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) page(
	w http.ResponseWriter,
	_ *http.Request,
) {
	sessions, e := s.service.EnrichedSessions(0, 0)
	errors.PanicOnError(e)
	limit := 30
	hasMore := len(sessions) > limit

	if hasMore {
		sessions = sessions[:limit]
	}

	var entries []gomponents.Node

	for _, e := range sessions {
		entries = append(entries, entry(e))
	}

	if hasMore {
		entries = append(entries, sentinel(limit))
	}

	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(
		layout(
			html.Div(
				html.Class("conversation-layout"),
				html.Div(
					html.Class("sidebar"),
					html.Input(
						html.Type("search"),
						html.Class("sidebar-filter"),
						gomponents.Attr("placeholder", "Filter..."),
						gomponents.Attr(
							"oninput",
							"filterSidebar(this.value)",
						),
					),
					html.Div(
						html.ID("sidebar-entries"),
						gomponents.Attr(
							"hx-get",
							"/conversations/sidebar",
						),
						gomponents.Attr(
							"hx-trigger",
							"session-edited from:body",
						),
						gomponents.Attr("hx-swap", "innerHTML"),
						gomponents.Group(entries),
					),
				),
				html.Div(
					html.Class("panel"),
					html.ID("panel"),
					html.P(
						html.Class("panel-placeholder"),
						gomponents.Text("Select a conversation"),
					),
				),
			),
		).Render(w),
	)
}
