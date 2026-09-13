package mechanic

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) appendItem(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.mutex.Lock()
	s.count++
	count := s.count
	s.mutex.Unlock()
	s.view.RenderFragment(w, html.Li(gomponents.Textf("item %d", count)))
}
