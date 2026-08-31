package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) plate(
	w http.ResponseWriter,
	_ *http.Request,
) {
	issues := s.worker.Issues()
	s.view.RenderLivePageWithSummary(
		w,
		constant.PlateTitle,
		constant.PlatePath,
		subscription.Query(
			constant.PlateEvent,
			constant.FavoritesEvent,
			constant.WatchedEvent,
			constant.SummaryEvent,
		),
		summary(issues),
		html.H3(gomponents.Text(constant.PlateTitle)),
		html.Div(
			gomponents.Attr("sse-swap", constant.PlateEvent),
			plateTable(issues),
		),
		html.H3(gomponents.Text(constant.FavoritesTitle)),
		html.Div(
			gomponents.Attr("sse-swap", constant.FavoritesEvent),
			pagesTable(s.worker.Favorites()),
		),
		html.H3(gomponents.Text(constant.WatchedTitle)),
		html.Div(
			gomponents.Attr("sse-swap", constant.WatchedEvent),
			pagesTable(s.worker.Watched()),
		),
	)
}
