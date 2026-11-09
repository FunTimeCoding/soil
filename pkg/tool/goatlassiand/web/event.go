package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web/plate_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleRequestServerSideEvent(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			subs := subscription.Parse(r)
			p := s.service.CurrentPlate()
			issues := p.Issues

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					webConstant.LayoutSummaryStrip,
					layout.SummaryStripContent(plate_page.Summary(issues)),
				)
			}

			if subs.Has(constant.NewestEvent) {
				layout.PushEvent(
					w,
					constant.NewestEvent,
					plate_page.IssuesSection(constant.NewestTitle, p.Newest),
				)
			}

			if subs.Has(constant.PlateEvent) {
				layout.PushEvent(
					w,
					constant.PlateEvent,
					plate_page.PlateSection(issues),
				)
			}

			if subs.Has(constant.WatchedIssuesEvent) {
				layout.PushEvent(
					w,
					constant.WatchedIssuesEvent,
					plate_page.IssuesSection(
						constant.WatchedIssuesTitle,
						p.WatchedIssues,
					),
				)
			}

			if subs.Has(constant.FavoritesEvent) {
				layout.PushEvent(
					w,
					constant.FavoritesEvent,
					plate_page.PagesSection(
						constant.FavoritesTitle,
						p.Favorites,
					),
				)
			}

			if subs.Has(constant.WatchedPagesEvent) {
				layout.PushEvent(
					w,
					constant.WatchedPagesEvent,
					plate_page.PagesSection(
						constant.WatchedPagesTitle,
						p.Watched,
					),
				)
			}

			f.Flush()
		},
	)
}
