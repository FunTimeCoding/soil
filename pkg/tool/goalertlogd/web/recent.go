package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/alert_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) recent(
	w http.ResponseWriter,
	r *http.Request,
) {
	given := window.Whole()

	if v := r.URL.Query().Get(constant.Start); v != "" {
		if t, e := time.Parse(constant.FilterFormat, v); e == nil {
			given.WithStart(t)
		}
	}

	if v := r.URL.Query().Get(constant.End); v != "" {
		if t, e := time.Parse(constant.FilterFormat, v); e == nil {
			given.WithEnd(t)
		}
	}

	bounded := s.service.RecentWindow(given)
	start, _ := bounded.Start()
	end, _ := bounded.End()
	records, e := s.service.RecentAlerts(bounded)
	errors.PanicOnError(e)
	table := alert_page.Recent(records)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, table)

		return
	}

	s.view.RenderPage(
		w,
		constant.RecentTitle,
		constant.RecentPath,
		html.H1(gomponents.Text("Recent Alerts")),
		alert_page.Filter(start, end),
		html.Div(html.ID(constant.RecentMark), table),
	)
}
