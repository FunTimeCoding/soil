package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) renderLogoutForm(
	w http.ResponseWriter,
	message string,
) {
	s.view.RenderPage(
		w,
		constant.SignOutTitle,
		constant.LogoutPath,
		html.P(gomponents.Text(message)),
		html.Form(
			html.Method("post"),
			html.Action(constant.LogoutPath),
			html.Button(
				html.Type("submit"),
				gomponents.Text(constant.SignOutTitle),
			),
		),
	)
}
