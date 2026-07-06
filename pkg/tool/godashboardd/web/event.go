package web

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEvent(
		s.service.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
		) {
			values := s.service.Values()

			for _, v := range s.board.Entries() {
				labels := rowLabels(v)

				if len(labels) == 0 {
					continue
				}

				layout.PushEvent(
					w,
					eventName(v.Label),
					rowSpans(labels, values[v.Label]),
				)
			}

			f.Flush()
		},
	)
}
