package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"net/http"
)

func (s *Server) sessionEditForm(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	found, okay, e := s.service.FindSession(identifier)

	if e == nil && !okay {
		e = not_found.New(constant.SessionKind, identifier)
	}

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPage(
		w,
		"Edit Session",
		constant.SessionsPath,
		session_page.EditForm(found),
	)
}
