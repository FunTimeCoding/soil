package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web/floor_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
					webConstant.LayoutSummaryStrip,
					layout.SummaryStripContent(floor_page.Summary(*snapshot)),
				)
			}

			if subs.Has(constant.FloorEvent) {
				layout.PushEvent(
					w,
					constant.FloorEvent,
					floor_page.Region(*snapshot, s.guestLinks(*snapshot)),
				)
			}

			f.Flush()
		},
	)
}
