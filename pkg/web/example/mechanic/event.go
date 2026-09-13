package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
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

			if subs.Has(constant.TickMark) {
				layout.PushEvent(w, constant.TickMark, s.tickCell())
			}

			if subs.Has(constant.PulseMark) {
				layout.PushEvent(w, constant.PulseMark, s.pulseCell())
			}

			f.Flush()
		},
	)
}
