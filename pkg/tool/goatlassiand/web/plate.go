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
			html.Div(
				extended.StreamSwap(constant.NewestEvent),
				issuesSection(constant.NewestTitle, s.worker.Newest()),
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
		html.Div(
			extended.StreamSwap(constant.PlateEvent),
			plateSection(issues),
		),
		html.Div(
			extended.StreamSwap(constant.WatchedIssuesEvent),
			issuesSection(
				constant.WatchedIssuesTitle,
				s.worker.WatchedIssues(),
			),
		),
		html.Div(
			extended.StreamSwap(constant.FavoritesEvent),
			pagesSection(constant.FavoritesTitle, s.worker.Favorites()),
		),
		html.Div(
			extended.StreamSwap(constant.WatchedPagesEvent),
			pagesSection(constant.WatchedPagesTitle, s.worker.Watched()),
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
