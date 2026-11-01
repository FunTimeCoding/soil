package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/session_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sessions(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows, e := s.service.AuthenticationSessions()
	errors.PanicOnError(e)
	current := currentSessionIdentifier(r)
	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		html.H1(gomponents.Textf("Sessions (%d)", len(rows))),
		gomponents.If(
			len(rows) > 1 && current != "",
			session_page.RevokeOthers(),
		),
		session_page.Table(rows, current),
	)
}
