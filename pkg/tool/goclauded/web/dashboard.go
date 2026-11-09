package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/usage_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	q *http.Request,
) {
	roster, e := s.rosterSection()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	activity, f := s.activity(q.Context(), nil)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	usage, _, g := s.usage.FindUsage()

	if g != nil {
		s.view.Answer(w, q, s.reporter, g)

		return
	}

	s.view.RenderSummaryLivePage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		subscription.Query(
			constant.Roster,
			constant.Activity,
			constant.EventSummary,
		),
		usage_page.Summary(usage),
		html.H3(gomponents.Text("Roster")),
		html.Div(extended.StreamSwap(constant.Roster), roster),
		html.H3(gomponents.Text("Recent Activity")),
		html.Div(extended.StreamSwap(constant.Activity), activity),
	)
}
