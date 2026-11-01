package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"net/http"
)

func (s *Server) reportDelete(
	w http.ResponseWriter,
	r *http.Request,
) {
	if e := s.service.DeleteReport(r.PathValue(constant.FileNameParameter)); e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	http.Redirect(w, r, constant.ReportsPath, http.StatusSeeOther)
}
