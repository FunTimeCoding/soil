package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) sessionDeleteAction(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	_, e := s.service.DeleteSession(
		identifier,
		s.service.DeleteHash(identifier),
	)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	http.Redirect(w, q, constant.SessionsPath, http.StatusSeeOther)
}
