package mechanic

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) subscribe(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.mutex.Lock()
	s.pulse++
	s.count++
	s.mutex.Unlock()
	s.notifier.Notify()
	s.view.RenderFragment(w, html.Span(gomponents.Text("pulsed")))
}
