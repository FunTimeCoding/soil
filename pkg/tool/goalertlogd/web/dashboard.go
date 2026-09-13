package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
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
		s.view.RenderFragment(w, s.topTable())

		return
	}

	s.view.RenderLivePageWithSummary(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		subscription.Query(constant.EventSummary, constant.EventTop),
		s.summaryItems(),
		html.H2(gomponents.Text("Top 25 Noisy Alerts (Last 7 Days)")),
		html.Div(
			html.ID(constant.TopMark),
			extended.StreamSwap(constant.EventTop),
			s.topTable(),
		),
	)
}
