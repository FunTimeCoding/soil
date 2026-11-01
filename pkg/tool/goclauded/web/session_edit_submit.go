package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"net/http"
)

func (s *Server) sessionEditSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	alias := q.FormValue(constant.Alias)
	description := q.FormValue(constant.Description)
	a := edit_session.New()
	a.Alias = &alias
	a.Description = &description

	if e := s.service.EditSession(identifier, a); e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	http.Redirect(w, q, session_page.Path(identifier), http.StatusSeeOther)
}
