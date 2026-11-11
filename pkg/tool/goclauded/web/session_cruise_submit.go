package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
)

func (s *Server) sessionCruiseSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	grace := q.FormValue(constant.GraceKind)

	if grace == constant.CruiseMinutes {
		grace = q.FormValue(constant.Grace)
	}

	refused := s.message.SetCruise(identifier, grace)

	if s.view.IsExtendedRequest(q) {
		s.cruiseFragment(w, q, identifier, refused)

		return
	}

	if refused != nil {
		form.IdentifierRedirect(
			w,
			q,
			session_page.Path(identifier),
			constant.CruiseAction,
			refused.Error(),
			nil,
		)

		return
	}

	http.Redirect(w, q, session_page.Path(identifier), http.StatusSeeOther)
}
