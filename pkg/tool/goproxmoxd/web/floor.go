package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) floor(
	w http.ResponseWriter,
	_ *http.Request,
) {
	f := s.worker.Floor()
	s.view.RenderLivePageWithSummary(
		w,
		constant.FloorTitle,
		constant.FloorPath,
		subscription.Query(constant.FloorEvent, constant.SummaryEvent),
		summary(*f),
		html.Div(
			gomponents.Attr("sse-swap", constant.FloorEvent),
			s.floorSections(*f),
		),
	)
}
