package mechanic

import "net/http"

func (s *Server) counterIncrement(
	w http.ResponseWriter,
	_ *http.Request,
) {
	s.mutex.Lock()
	s.count++
	s.mutex.Unlock()
	s.view.RenderFragment(w, s.counterCell())
}
