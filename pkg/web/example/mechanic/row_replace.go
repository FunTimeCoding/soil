package mechanic

import "net/http"

func (s *Server) rowReplace(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderFragment(w, rowCell(true))
}
