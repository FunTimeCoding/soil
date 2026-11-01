package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"net/http"
)

func (s *Server) reportDownload(
	w http.ResponseWriter,
	r *http.Request,
) {
	path, e := s.service.ReportPath(r.PathValue(constant.FileNameParameter))

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	http.ServeFile(w, r, path)
}
