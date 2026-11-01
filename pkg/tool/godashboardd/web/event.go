package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web/board_page"
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
			p := board_page.New(s.board, s.service.Values())

			for _, v := range s.board.Entries() {
				rows := p.Rows(v)

				if rows == nil {
					continue
				}

				layout.PushEvent(w, board_page.EventName(v.Label), rows)
			}

			f.Flush()
		},
	)
}
