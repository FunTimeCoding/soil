package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/player_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) players(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows := s.service.PlayerAttendance()
	s.view.RenderPage(
		w,
		constant.PlayersTitle,
		constant.PlayersPath,
		html.H1(gomponents.Textf("Player Attendance (%d)", len(rows))),
		player_page.List(rows),
	)
}
