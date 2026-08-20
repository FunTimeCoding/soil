package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderLivePageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		subscription.Query(
			constant.Roster,
			constant.Activity,
			constant.EventSummary,
		),
		s.usageSummary(),
		html.H3(gomponents.Text("Roster")),
		html.Div(
			gomponents.Attr("sse-swap", constant.Roster),
			s.rosterSection(),
		),
		html.H3(gomponents.Text("Recent Activity")),
		html.Div(
			gomponents.Attr("sse-swap", constant.Activity),
			s.activitySection(nil),
		),
	)
}
