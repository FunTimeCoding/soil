package mechanic

import (
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) quiet(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.view.RenderFragment(w, gomponents.Text("ignored"))
}
