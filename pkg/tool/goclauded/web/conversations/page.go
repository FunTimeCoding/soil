package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) page(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(
		layout(
			html.Div(
				html.Class("conversation-layout"),
				html.Div(
					html.Class("sidebar"),
					searchForm(),
					html.Div(
						html.ID("sidebar-entries"),
						extended.Get("/conversations/sidebar"),
						extended.Trigger("session-edited from:body"),
						extended.Swap("innerHTML"),
						gomponents.Group(s.sidebarNodes(0)),
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
