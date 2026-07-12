package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
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
			layout.PushEvent(w, constant.Seeds, s.seedTable())
			layout.PushEvent(w, constant.SeedsRecent, s.recentSeedTable())
			f.Flush()
		},
	)
}
