package web

import "net/http"

func (s *Server) signIn(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.authorization.SignIn(w, r)
}
