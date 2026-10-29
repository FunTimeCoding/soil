package server

import "net/http"

func (s *Server) logoutGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	row, e := s.authenticatedSession(r)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	if row == nil {
		s.renderLogoutForm(w, "You are already signed out.")

		return
	}

	s.renderLogoutForm(w, "Sign out of every service?")
}
