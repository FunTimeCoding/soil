package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
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
	content := []gomponents.Node{
		html.H1(gomponents.Textf("Sessions (%d)", len(rows))),
	}

	if len(rows) > 1 && current != "" {
		content = append(
			content,
			html.Form(
				html.Method("post"),
				html.Action(constant.SessionsRevokeOthersPath),
				html.Button(
					html.Type("submit"),
					gomponents.Attr("class", "secondary"),
					gomponents.Text("Revoke all other sessions"),
				),
			),
		)
	}

	content = append(content, sessionsTable(rows, current))
	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		content...,
	)
}
