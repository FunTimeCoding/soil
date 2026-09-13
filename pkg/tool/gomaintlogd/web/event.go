package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			subs := subscription.Parse(r)

			if subs.Has(constant.EventSummary) {
				layout.PushEvent(
					w,
					webConstant.LayoutSummaryStrip,
					layout.SummaryStripContent(s.summaryItems()),
				)
			}

			if subs.Has(constant.EventRecent) {
				layout.PushEvent(w, constant.EventRecent, s.recentTable())
			}

			f.Flush()
		},
	)
}
