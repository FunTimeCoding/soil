package web

import "net/http"

func (s *Server) signOut(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.authorization.SignOut(w, r)
}
