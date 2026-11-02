package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web/plate_page"
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
	p := s.service.CurrentPlate()
	issues := p.Issues
	var events []string
	var content []gomponents.Node

	if s.service.HasProject() {
		events = append(events, constant.NewestEvent)
		content = append(
			content,
			html.Div(
				extended.StreamSwap(constant.NewestEvent),
				plate_page.IssuesSection(constant.NewestTitle, p.Newest),
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
			plate_page.PlateSection(issues),
		),
		html.Div(
			extended.StreamSwap(constant.WatchedIssuesEvent),
			plate_page.IssuesSection(
				constant.WatchedIssuesTitle,
				p.WatchedIssues,
			),
		),
		html.Div(
			extended.StreamSwap(constant.FavoritesEvent),
			plate_page.PagesSection(constant.FavoritesTitle, p.Favorites),
		),
		html.Div(
			extended.StreamSwap(constant.WatchedPagesEvent),
			plate_page.PagesSection(constant.WatchedPagesTitle, p.Watched),
		),
	)
	s.view.RenderLivePageWithSummary(
		w,
		constant.PlateTitle,
		constant.PlatePath,
		subscription.Query(events...),
		plate_page.Summary(issues),
		content...,
	)
}
