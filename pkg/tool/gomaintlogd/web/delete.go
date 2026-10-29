package web

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := identifierFromQuery(r)

	if e == nil {
		e = s.service.Delete(identifier)
	}

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderStatus(w, "Deleted entry", webConstant.LayoutStatusSuccess)
}
