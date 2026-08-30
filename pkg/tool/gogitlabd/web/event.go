package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
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
			entries := s.worker.Entries()

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					web.LayoutSummaryStrip,
					layout.SummaryStripContent(summary(entries)),
				)
			}

			if subs.Has(constant.BoardEvent) {
				layout.PushEvent(w, constant.BoardEvent, s.boardTable(entries))
			}

			f.Flush()
		},
	)
}
