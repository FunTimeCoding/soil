package mechanic

import "net/http"

func (s *Server) counter(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderFragment(w, s.counterCell())
}
