package web

import "net/http"

func (s *Server) require(next http.HandlerFunc) http.HandlerFunc {
	if s.authorization == nil {
		return next
	}

	return s.authorization.Require(next)
}
