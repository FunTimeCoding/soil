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
		d := constant.DeliveryQueued

		if r.FormValue(constant.Immediate) != "" {
			d = constant.DeliveryImmediate
		}

		_, e := s.service.SendPulse(identifier, "", body, d)
		errors.PanicOnError(e)
	}

	http.Redirect(
		w,
		r,
		fmt.Sprintf("/sessions/%s", identifier),
		http.StatusSeeOther,
	)
}
