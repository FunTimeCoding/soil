package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
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

	s.view.RenderLivePageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		subscription.Query(constant.EventSummary, constant.EventRecent),
		s.summaryItems(),
		html.H2(gomponents.Text("Recent Entries")),
		html.Div(
			html.ID(constant.RecentMark),
			extended.StreamSwap(constant.EventRecent),
			s.recentTable(),
		),
	)
}
