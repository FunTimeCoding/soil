package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/counter_page"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/seed_page"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.service.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
			q *http.Request,
		) {
			layout.PushEvent(
				w,
				constant.Seeds,
				seed_page.Priority(s.service.Seeds()),
			)
			layout.PushEvent(
				w,
				constant.SeedsRecent,
				seed_page.Recent(s.service.RecentSeeds()),
			)
			layout.PushEvent(
				w,
				constant.SessionsEvent,
				session_page.List(s.service.Sessions()),
			)
			layout.PushEvent(
				w,
				constant.CounterEvent,
				counter_page.Table(s.service.Tallies()),
			)

			if name := q.URL.Query().Get(constant.SessionParameter); name != "" {
				pushDecisions(w, s.service.Decisions(name))
			}

			f.Flush()
		},
	)
}
