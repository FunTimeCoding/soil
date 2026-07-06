package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, s.recentTable())

		return
	}

	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H1(gomponents.Text(constant.DashboardTitle)),
		html.Div(
			html.Class("summary-cards"),
			html.Article(
				html.Header(gomponents.Text("Total Entries")),
				html.P(gomponents.Textf("%d", s.store.Count())),
			),
		),
		html.H2(gomponents.Text("Recent Entries")),
		html.Div(
			html.ID("recent-table"),
			htmx.Get(constant.DashboardPath),
			htmx.Trigger("every 60s"),
			htmx.Swap("innerHTML"),
			s.recentTable(),
		),
	)
}
