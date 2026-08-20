package web

import "net/http"

func (s *Server) require(next http.HandlerFunc) http.HandlerFunc {
	return s.authorization.Require(next)
}
