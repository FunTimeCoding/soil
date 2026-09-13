package mechanic

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) poll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.mutex.Lock()
	s.pulse++
	pulse := s.pulse
	s.mutex.Unlock()
	s.view.RenderFragment(w, html.Span(gomponents.Textf("poll %d", pulse)))
}
