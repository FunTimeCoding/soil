package web

import "net/http"

func (s *Server) callback(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.authorization.Callback(w, r)
}
