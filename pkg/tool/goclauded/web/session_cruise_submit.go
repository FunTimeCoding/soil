package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) sessionCruiseSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(r.ParseForm())
	value := r.FormValue(constant.Mode)
	mode := constant.Cruise(value)
	var pace time.Duration

	if minutes, e := strconv.Atoi(value); e == nil {
		mode = constant.CruisePaced
		pace = time.Duration(minutes) * time.Minute
	}

	if e := s.service.SetCruise(identifier, mode, pace); e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	http.Redirect(
		w,
		r,
		fmt.Sprintf("/sessions/%s", identifier),
		http.StatusSeeOther,
	)
}
