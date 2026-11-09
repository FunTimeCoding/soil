package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) sessionCruiseSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	mode := constant.Cruise(q.FormValue(constant.Mode))
	var pace time.Duration

	if minutes, e := strconv.ParseFloat(q.FormValue(constant.Pace), 64); e == nil {
		pace = time.Duration(minutes * float64(time.Minute))
	}

	refused := s.message.SetCruise(identifier, mode, pace)

	if s.view.IsExtendedRequest(q) {
		s.cruiseFragment(w, q, identifier, mode, pace, refused)

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
