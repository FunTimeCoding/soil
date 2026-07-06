package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) playerDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	account := r.PathValue("account")
	rows := s.store.PlayerDetail(account)
	s.view.RenderPage(
		w,
		account,
		constant.PlayersPath,
		html.H1(gomponents.Text(account)),
		playerDetailTable(rows),
	)
}
