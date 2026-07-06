package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) players(
	w http.ResponseWriter,
	r *http.Request,
) {
	since := time.Now().AddDate(0, -1, 0)
	rows := s.store.PlayerAttendance(since)
	s.view.RenderPage(
		w,
		constant.PlayersTitle,
		constant.PlayersPath,
		html.H1(gomponents.Textf("Player Attendance (%d)", len(rows))),
		playersTable(rows),
	)
}
