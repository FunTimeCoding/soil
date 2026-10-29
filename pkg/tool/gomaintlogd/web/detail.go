package web

import "net/http"

func (s *Server) detail(
	w http.ResponseWriter,
	r *http.Request,
) {
	e, f := s.entryFromQuery(r)

	if f != nil {
		s.view.Answer(w, r, s.reporter, f)

		return
	}

	s.view.RenderFragment(w, detailRow(e))
}
