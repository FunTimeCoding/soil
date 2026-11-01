package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/usage_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
			q *http.Request,
		) {
			subs := subscription.Parse(q)

			if subs.Has(constant.EventSummary) {
				if items := usage_page.Summary(s.service.Usage()); len(items) > 0 {
					layout.PushEvent(
						w,
						webConstant.LayoutSummaryStrip,
						layout.SummaryStripContent(items),
					)
				}
			}

			if subs.Has(constant.Usage) {
				layout.PushEvent(
					w,
					constant.UsageChart,
					usage_page.Chart(s.service.UsageWindow()),
				)
			}

			if subs.Has(constant.Roster) {
				s.push(w, constant.Roster, s.roster)
			}

			if subs.Has(constant.Activity) {
				s.push(
					w,
					constant.Activity,
					func() (gomponents.Node, error) {
						return s.activity(q.Context(), nil)
					},
				)
			}

			identifier := q.URL.Query().Get(constant.EventSession)

			if subs.Has(constant.Pulse) && identifier != "" {
				pulses, e := s.service.PulsesBySession(identifier)

				if e != nil {
					s.reporter.CaptureException(e)
				} else {
					layout.PushEvent(
						w,
						constant.Pulse,
						session_page.Pulses(pulses),
					)
				}
			}

			f.Flush()
		},
	)
}
