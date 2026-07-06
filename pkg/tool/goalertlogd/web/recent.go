package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) recent(
	w http.ResponseWriter,
	r *http.Request,
) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v := r.URL.Query().Get(constant.Start); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			start = t
		}
	}

	if v := r.URL.Query().Get(constant.End); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			end = t
		}
	}

	table := recentTable(s.store.MustByTimeRange(start, end))

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, table)

		return
	}

	s.view.RenderPage(
		w,
		constant.RecentTitle,
		constant.RecentPath,
		html.H1(gomponents.Text("Recent Alerts")),
		html.Form(
			html.Class("filter-form"),
			htmx.Get(constant.RecentPath),
			htmx.Target("#recent-table"),
			htmx.Swap("innerHTML"),
			html.Div(
				html.Class("grid"),
				html.Label(
					gomponents.Text("Start"),
					html.Input(
						html.Type("datetime-local"),
						html.Name(constant.Start),
						html.Value(start.Format("2006-01-02T15:04")),
					),
				),
				html.Label(
					gomponents.Text("End"),
					html.Input(
						html.Type("datetime-local"),
						html.Name(constant.End),
						html.Value(end.Format("2006-01-02T15:04")),
					),
				),
				html.Label(
					gomponents.Raw("&nbsp;"),
					html.Button(
						html.Type("submit"),
						gomponents.Text("Filter"),
					),
				),
			),
		),
		html.Div(html.ID("recent-table"), table),
	)
}
