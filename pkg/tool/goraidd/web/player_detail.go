package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web/player_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) playerDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	account := r.PathValue(constant.AccountParameter)
	rows := s.service.PlayerDetail(account)
	s.view.RenderPage(
		w,
		account,
		constant.PlayersPath,
		html.H1(gomponents.Text(account)),
		player_page.Detail(rows),
	)
}
