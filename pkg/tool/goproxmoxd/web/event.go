package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
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
			snapshot := s.worker.Floor()

			if subs.Has(constant.SummaryEvent) {
				layout.PushEvent(
					w,
					web.LayoutSummaryStrip,
					layout.SummaryStripContent(summary(*snapshot)),
				)
			}

			if subs.Has(constant.FloorEvent) {
				layout.PushEvent(
					w,
					constant.FloorEvent,
					s.floorSections(*snapshot),
				)
			}

			f.Flush()
		},
	)
}
