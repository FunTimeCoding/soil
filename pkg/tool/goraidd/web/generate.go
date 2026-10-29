package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"net/http"
)

func (s *Server) generate(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())

	if e := s.service.Generate(r.Context(), r.Form["fileNames"]); e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	http.Redirect(w, r, constant.ReportsPath, http.StatusSeeOther)
}
