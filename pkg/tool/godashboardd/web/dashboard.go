package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web/board_page"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	_ *http.Request,
) {
	p := board_page.New(s.board, s.service.Values())
	s.view.RenderPage(
		w,
		constant.DashboardTitle,
		constant.DashboardPath,
		p.Top(),
		p.Tail(),
	)
}
