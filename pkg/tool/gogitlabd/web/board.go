package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) board(
	w http.ResponseWriter,
	_ *http.Request,
) {
	entries := s.worker.Entries()
	s.view.RenderLivePageWithSummary(
		w,
		constant.BoardTitle,
		constant.BoardPath,
		subscription.Query(constant.BoardEvent, constant.SummaryEvent),
		summary(entries),
		html.Div(
			gomponents.Attr("sse-swap", constant.BoardEvent),
			s.boardTable(entries),
		),
	)
}
