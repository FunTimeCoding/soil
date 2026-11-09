package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/board_page"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) board(
	w http.ResponseWriter,
	_ *http.Request,
) {
	entries := s.service.Entries()
	multiple := s.multiple()
	s.view.RenderSummaryLivePage(
		w,
		constant.BoardTitle,
		constant.BoardPath,
		subscription.Query(
			constant.RequestEvent,
			constant.BoardEvent,
			constant.SummaryEvent,
		),
		board_page.Summary(entries),
		html.Div(
			extended.StreamSwap(constant.RequestEvent),
			board_page.Requests(s.service.Requests(), multiple),
		),
		html.H3(gomponents.Text(constant.BoardTitle)),
		html.Div(
			extended.StreamSwap(constant.BoardEvent),
			board_page.Table(entries, multiple),
		),
	)
}
