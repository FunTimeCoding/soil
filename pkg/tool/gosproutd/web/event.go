package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.service.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			layout.PushEvent(w, constant.Seeds, s.seedTable())
			layout.PushEvent(w, constant.SeedsRecent, s.recentSeedTable())
			layout.PushEvent(w, constant.SessionsEvent, s.sessionTable())
			layout.PushEvent(w, constant.CounterEvent, s.tallyTable())

			if name := r.URL.Query().Get(
				constant.SessionParameter,
			); name != "" {
				layout.PushEvent(
					w,
					constant.DecisionEvent,
					s.decisionTable(name),
				)
			}

			f.Flush()
		},
	)
}
