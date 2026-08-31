package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.worker.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			subs := subscription.Parse(r)
			issues := s.worker.Issues()

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					web.LayoutSummaryStrip,
					layout.SummaryStripContent(summary(issues)),
				)
			}

			if subs.Has(constant.PlateEvent) {
				layout.PushEvent(w, constant.PlateEvent, plateTable(issues))
			}

			if subs.Has(constant.FavoritesEvent) {
				layout.PushEvent(
					w,
					constant.FavoritesEvent,
					pagesTable(s.worker.Favorites()),
				)
			}

			if subs.Has(constant.WatchedEvent) {
				layout.PushEvent(
					w,
					constant.WatchedEvent,
					pagesTable(s.worker.Watched()),
				)
			}

			f.Flush()
		},
	)
}
