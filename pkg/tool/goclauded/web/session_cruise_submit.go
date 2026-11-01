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
	value := q.FormValue(constant.Mode)
	mode := constant.Cruise(value)
	var pace time.Duration

	if minutes, e := strconv.Atoi(value); e == nil {
		mode = constant.CruisePaced
		pace = time.Duration(minutes) * time.Minute
	}

	if e := s.service.SetCruise(identifier, mode, pace); e != nil {
		form.RedirectFor(
			w,
			q,
			session_page.Path(identifier),
			constant.CruiseAction,
			e.Error(),
			nil,
		)

		return
	}

	http.Redirect(w, q, session_page.Path(identifier), http.StatusSeeOther)
}
