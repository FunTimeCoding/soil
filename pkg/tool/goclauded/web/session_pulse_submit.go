package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) sessionPulseSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(r.ParseForm())
	body := r.FormValue(constant.Body)

	if body != "" {
		_, e := s.service.SendPulse(
			identifier,
			"",
			body,
			r.FormValue(constant.Immediate) != "",
		)
		errors.PanicOnError(e)
	}

	http.Redirect(
		w,
		r,
		fmt.Sprintf("/sessions/%s", identifier),
		http.StatusSeeOther,
	)
}
