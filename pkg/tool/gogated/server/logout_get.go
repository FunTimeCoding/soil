package server

import "net/http"

func (s *Server) logoutGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.authenticatedSession(r) == nil {
		s.renderLogoutForm(w, "You are already signed out.")

		return
	}

	s.renderLogoutForm(w, "Sign out of every service?")
}
