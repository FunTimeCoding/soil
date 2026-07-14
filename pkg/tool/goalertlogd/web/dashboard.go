package web

import (
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
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
		s.view.RenderFragment(w, s.topTable())

		return
	}

	lastPoll := s.worker.LastPoll()
	lastPollText := "never"

	if !lastPoll.IsZero() {
		lastPollText = time.FormatCompact(lastPoll)
	}

	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		html.H1(gomponents.Text(constant.DashboardTitle)),
		html.Div(
			html.Class("summary-cards"),
			html.Article(
				html.Header(gomponents.Text("Total Records")),
				html.P(gomponents.Textf("%d", s.store.MustCount())),
			),
			html.Article(
				html.Header(gomponents.Text("Currently Firing")),
				html.P(
					gomponents.Textf("%d", s.worker.FiringCount()),
				),
			),
			html.Article(
				html.Header(gomponents.Text("Last Poll")),
				html.P(gomponents.Text(lastPollText)),
			),
		),
		html.H2(gomponents.Text("Top 25 Noisy Alerts (Last 7 Days)")),
		html.Div(
			html.ID("top-table"),
			htmx.Get(constant.DashboardPath),
			htmx.Trigger("every 60s"),
			htmx.Swap("innerHTML"),
			s.topTable(),
		),
	)
}
