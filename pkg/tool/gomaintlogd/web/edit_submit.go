package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"net/http"
	"time"
)

func (s *Server) editSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	errors.PanicOnError(r.ParseForm())
	e.Action = r.FormValue(constant.Action)
	e.User = r.FormValue(constant.User)
	e.System = r.FormValue(constant.System)
	e.Service = r.FormValue(constant.Service)
	e.Description = r.FormValue(constant.Description)

	if v := r.FormValue(constant.Timestamp); v != "" {
		if t, f := time.Parse("2006-01-02T15:04", v); f == nil {
			e.Timestamp = t
		}
	}

	errors.PanicOnError(s.store.Update(e))
	s.view.RenderFragment(w, detailRow(e))
}
