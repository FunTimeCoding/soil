package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
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
			subs := parseSubscribe(r)

			if subs.Has(constant.EventSummary) {
				if items := s.usageSummary(); len(items) > 0 {
					layout.PushEvent(
						w,
						layout.SummaryStrip,
						layout.SummaryStripContent(items),
					)
				}
			}

			if subs.Has(constant.Roster) {
				layout.PushEvent(
					w,
					constant.Roster,
					s.rosterSection(),
				)
			}

			if subs.Has(constant.Activity) {
				layout.PushEvent(
					w,
					constant.Activity,
					s.activitySection(nil),
				)
			}

			if subs.Has(constant.Pulse) {
				identifier := r.URL.Query().Get(constant.EventSession)

				if identifier != "" {
					layout.PushEvent(
						w,
						constant.Pulse,
						s.pulseSection(identifier),
					)
				}
			}

			f.Flush()
		},
	)
}
