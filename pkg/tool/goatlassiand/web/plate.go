package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
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
	var events []string
	var content []gomponents.Node

	if s.worker.HasProject() {
		events = append(events, constant.NewestEvent)
		content = append(
			content,
			html.H3(gomponents.Text(constant.NewestTitle)),
			html.Div(
				extended.StreamSwap(constant.NewestEvent),
				issuesTable(s.worker.Newest(), constant.NewestEmpty),
			),
		)
	}

	events = append(
		events,
		constant.PlateEvent,
		constant.WatchedIssuesEvent,
		constant.FavoritesEvent,
		constant.WatchedPagesEvent,
		constant.SummaryEvent,
	)
	content = append(
		content,
		html.H3(gomponents.Text(constant.PlateTitle)),
		html.Div(
			extended.StreamSwap(constant.PlateEvent),
			issuesTable(issues, constant.PlateEmpty),
		),
		html.H3(gomponents.Text(constant.WatchedIssuesTitle)),
		html.Div(
			extended.StreamSwap(constant.WatchedIssuesEvent),
			issuesTable(s.worker.WatchedIssues(), constant.WatchedIssuesEmpty),
		),
		html.H3(gomponents.Text(constant.FavoritesTitle)),
		html.Div(
			extended.StreamSwap(constant.FavoritesEvent),
			pagesTable(s.worker.Favorites()),
		),
		html.H3(gomponents.Text(constant.WatchedPagesTitle)),
		html.Div(
			extended.StreamSwap(constant.WatchedPagesEvent),
			pagesTable(s.worker.Watched()),
		),
	)
	s.view.RenderLivePageWithSummary(
		w,
		constant.PlateTitle,
		constant.PlatePath,
		subscription.Query(events...),
		summary(issues),
		content...,
	)
}
