package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"net/http"
)

func (s *Server) publishSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	if _, _, e := s.service.Publish(); e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	http.Redirect(w, r, constant.DashboardPath, http.StatusSeeOther)
}
